package main

import (
	"fmt"
	"time"

	"github.com/forward-step/go_progress/progress"
)

func main() {
	factory := progress.New(
		progress.OptionPrintRecord(func(r progress.Record) string {
			return fmt.Sprintf("%.2f it/s  %s", r.CurrentRate(), progress.DefaultPrintRecord(r))
		}),
	)
	p := factory.Add(100)

	for p.Add(1) {
		time.Sleep(time.Millisecond * 100)
	}

	<-factory.Done
}
