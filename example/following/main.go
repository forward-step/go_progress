package main

import (
	"fmt"
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
	fmt.Println("done1")

	f2 := progress.New()
	p1 := f2.Add(100)
	p2 := f2.Add(100)
	go func() {
		for p1.Add(8) {
			time.Sleep(time.Millisecond * 100)
		}
	}()
	go func() {
		for p2.Add(10) {
			time.Sleep(time.Millisecond * 100)
		}
	}()
	<-f2.Done
	fmt.Println("done2")
}
