package progress

import (
	"log"
	"sync"
	"time"
)

var TermWidth int

type Result struct {
	Length        uint8
	FailNumber    uint8
	SuccessNumber uint8
}
type FinshFn func(Result)
type FailStrategy uint8

const (
	FAIL_STRATEGY_SHUTDOWN = iota // 默认 ; 一个失败全部失败
	FAIL_STRATEGY_IGNORE          // 忽略 ; 其他继续执行
)

type ProgressFactory struct {
	sync.RWMutex
	arr           []*Progress   // 进度条数组
	debug         bool          // 调试模式
	onFinsh       FinshFn       // 完成函数
	updateCh      chan struct{} // 立即更新 ; 请使用`update`函数
	Done          chan Result   // 全部结束
	line          int           // 上一次打印行数
	historyLength uint          // 历史记录的长度
	failStrategy  FailStrategy  // 进度条失败后执行的策略
	printConfig   printConfig
}

func init() {
	var err error
	TermWidth, err = termWidth()
	if err != nil {
		log.Fatalln(err)
	}
}

func New(options ...Option) *ProgressFactory {
	var pf = &ProgressFactory{
		printConfig: defaultPrintConfig(),
		updateCh:    make(chan struct{}),
		Done:        make(chan Result),
	}

	for _, fn := range options {
		fn(pf)
	}

	pf.historyLength = max(2, pf.historyLength)

	go func() {
		pf.Lock()
		duration := pf.printConfig.refreshRate
		pf.Unlock()
		timer := time.NewTimer(duration)
		for {
			timer.Reset(duration)
			select {
			case <-pf.updateCh:
				pf.refreshConsole()
			case <-timer.C:
				pf.refreshConsole()
			}
		}
	}()

	return pf
}

// if blocked then lost
func (pf *ProgressFactory) update() {
	select {
	case pf.updateCh <- struct{}{}:
	default:
	}
}

func (pf *ProgressFactory) Add(total int64) *Progress {
	pf.Lock()
	defer pf.Unlock()
	p := &Progress{
		total:   total,
		start:   time.Now(),
		factory: pf,
		status:  STATUS_LOADING,
	}
	if pf.historyLength > 0 {
		p.history = make(History, pf.historyLength)
		p.history = append(p.history, HistoryItem{
			Time: p.start,
			Incr: 0,
		})
	}
	pf.arr = append(pf.arr, p)
	pf.update()
	return p
}

func (pf *ProgressFactory) finish(r Result) {
	pf.updateCh = nil
	// close(f.updateCh)
	if pf.onFinsh != nil {
		pf.onFinsh(r)
	}
}

func (pf *ProgressFactory) fail() {
	for _, p := range pf.arr {
		p.status = STATUS_FAIL
	}
	pf.update()
}
