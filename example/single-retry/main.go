package main

import (
	"time"

	"github.com/forward-step/go_progress/progress"
)

func main() {
	factory := progress.New()
	p := factory.Add(100)

	// mock retry
	go func() {
		<-time.After(time.Second * 3)
		p.Reset()
		for p.Add(1) {
			time.Sleep(time.Millisecond * 100)
		}
	}()

	for p.Add(1) {
		time.Sleep(time.Millisecond * 100)
	}

	<-factory.Done
}
