package progress

import (
	"time"
)

type RecordWithoutRate struct {
	Index       int
	Percent     float64
	Time        string
	ElapsedTime time.Duration
	Current     int64
	Total       int64
	Status      StatusValue
	Data        any
}

type HistoryItem struct {
	Time time.Time // 当前时间点
	Incr int64     // 添加大小
}

type History []HistoryItem

type Record struct {
	RecordWithoutRate
	Rate         string
	LoadingIndex int
	History      History
}

type RecordSummary struct {
	Result
	Record
}

func (p *Progress) getRecord(index int) Record {
	var f = p.factory
	percent := float64(p.current) / float64(p.total)

	var recordWithoutRate = RecordWithoutRate{
		Index:       index,
		Percent:     percent * 100,
		Time:        elapsedTimeByDuration(p.elapsedTime),
		ElapsedTime: p.elapsedTime,
		Current:     p.current,
		Total:       p.total,
		Status:      p.status,
		Data:        p.data,
	}

	var rate = recordWithoutRate.print(f.printConfig)

	return Record{
		RecordWithoutRate: recordWithoutRate,
		Rate:              rate,
		History:           p.history,
	}
}

// computed current rate
func (r *Record) CurrentRate() (rate float64) {
	length := len(r.History)
	if r.History != nil && length >= 2 {
		it1 := r.History[length-1]
		it2 := r.History[length-2]
		duration := float64(it1.Time.Sub(it2.Time)) / float64(time.Second)
		size := float64(it1.Incr)
		rate = size / duration
	}
	return
}

// computed average rate
func (r *Record) AverageRate() float64 {
	duration := float64(r.ElapsedTime) / float64(time.Second)
	return float64(r.Current) / duration
}
