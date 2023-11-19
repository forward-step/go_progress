package main

import (
	"time"

	"github.com/forward-step/go_progress/progress"
)

func main() {
	factory := progress.New(
		progress.OptionPrintFill('█'),
		progress.OptionPrintEmpty(' '),
		progress.OptionPrintHead(0),
		progress.OptionPrintLeft('['),
		progress.OptionPrintRight(']'),
	)
	p := factory.Add(100)

	for p.Add(1) {
		time.Sleep(time.Millisecond * 100)
	}

	<-factory.Done
}
