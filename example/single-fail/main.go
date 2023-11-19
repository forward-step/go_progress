package main

import (
	"sync"
	"time"

	"github.com/forward-step/go_progress/progress"
)

var wg sync.WaitGroup

func main() {
	factory := progress.New()
	p := factory.Add(100)

	// mock fail
	go func() {
		<-time.After(time.Second * 3)
		p.Fail()
	}()

	for p.Add(1) {
		time.Sleep(time.Millisecond * 100)
	}

	<-factory.Done
}
