package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/forward-step/go_progress/progress"
)

func main() {
	f1 := progress.New(
		progress.OptionFailStrategy(progress.FAIL_STRATEGY_IGNORE),
		progress.OptionPrintAppendRecord(progress.DefaultPrintAppendRecord),
		progress.OptionPrintRecord(func(r progress.Record) string {
			// only print progress of loading status
			if r.LoadingIndex != -1 && r.LoadingIndex < 3 {
				return fmt.Sprintf("index%d: %s", r.Index, progress.DefaultPrintRecord(r))
			}
			return ""
		}),
	)
	for i := 0; i < 100; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		go func() {
			p := f1.Add(100)
			add := r.Int63n(10) + 1 // [1, 10]
			for p.Add(add) {
				time.Sleep(time.Millisecond * 100)
			}
		}()
	}

	<-f1.Done
	defer close(f1.Done)
}
