package progress

import (
	"sync"
	"time"
)

type StatusValue uint8

const (
	STATUS_LOADING StatusValue = iota
	STATUS_SUCCESS
	STATUS_FAIL
)

type Progress struct {
	sync.Mutex
	status      StatusValue
	history     History
	current     int64
	total       int64
	start       time.Time
	elapsedTime time.Duration
	factory     *ProgressFactory
	data        any
}

func (p *Progress) Mount(data any) *Progress {
	p.data = data
	return p
}

func (p *Progress) Add(i int64) bool {
	return p.Set(p.current + i)
}

func (p *Progress) Set(i int64) bool {
	p.Lock()
	defer p.Unlock()

	// reject set
	if p.isEnd() {
		return false
	}

	// set
	var now = time.Now()
	if p.history != nil {
		var limitLength = p.factory.historyLength - 1
		var historyLength = len(p.history)
		var startIndex = max(0, historyLength-int(limitLength))
		p.history = append(p.history[startIndex:], HistoryItem{
			Time: now,
			Incr: i - p.current,
		})
	}
	p.current = min(i, p.total)
	p.elapsedTime = now.Sub(p.start)
	if p.current >= p.total {
		p.status = STATUS_SUCCESS
	}
	p.factory.update()

	return !p.isEnd()
}

func (p *Progress) Reset() {
	p.Lock()
	defer p.Unlock()

	p.history = p.history[:0]
	p.current = 0
	p.status = STATUS_LOADING
	p.factory.update()
}

func (p *Progress) Success() {
	p.Lock()
	defer p.Unlock()

	p.current = p.total
	p.status = STATUS_SUCCESS
	p.factory.update()
}

func (p *Progress) Fail() {
	p.Lock()
	defer p.Unlock()

	switch p.factory.failStrategy {
	case FAIL_STRATEGY_SHUTDOWN:
		p.factory.fail()
	case FAIL_STRATEGY_IGNORE:
		p.status = STATUS_FAIL
		p.factory.update()
	}
}

func (p *Progress) isEnd() bool {
	return p.status == STATUS_SUCCESS || p.status == STATUS_FAIL
}

func (p *Progress) isDone() bool {
	return p.current >= p.total
}
