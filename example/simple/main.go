package main

import (
	"time"

	"github.com/forward-step/go_progress/progress"
)

func main() {
	f1 := progress.New()
	p := f1.Add(100)

	for p.Add(10) {
		time.Sleep(time.Millisecond * 100)
	}

	<-f1.Done
	defer close(f1.Done)
}
