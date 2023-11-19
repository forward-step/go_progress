package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/forward-step/go_progress/loading"
	"github.com/forward-step/go_progress/tool"
)

func main() {
	l := loading.New(0)
	length := 10
	for i := length; i >= 0; i-- {
		if i != length {
			tool.ClearLine(os.Stdout, 2)
		}
		fmt.Printf("%s\n%s", l.Next(), strings.Repeat("=", i))
		time.Sleep(time.Millisecond * 200)
	}

}
