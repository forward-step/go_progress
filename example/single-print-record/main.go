package main

import (
	"fmt"
	"time"

	"github.com/forward-step/go_progress/loading"
	"github.com/forward-step/go_progress/progress"
)

func main() {
	l := loading.New(0)
	factory := progress.New(
		progress.OptionPrintRecord(func(r progress.Record) string {
			return fmt.Sprintf("%s    %s", l.Next(), progress.DefaultPrintRecord(r))
		}),
	)
	p := factory.Add(100)

	for p.Add(10) {
		time.Sleep(time.Millisecond * 200)
	}

	<-factory.Done
}
