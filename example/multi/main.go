package main

import (
	"time"

	"github.com/forward-step/go_progress/progress"
)

func main() {
	factory := progress.New()

	p1 := factory.Add(100)
	go func() {
		for p1.Add(1) {
			time.Sleep(100 * time.Millisecond)
		}
	}()

	p2 := factory.Add(300)
	go func() {
		for p2.Add(2) {
			time.Sleep(100 * time.Millisecond)
		}
	}()

	p3 := factory.Add(200)
	go func() {
		for p3.Add(3) {
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// wait
	<-factory.Done
}
