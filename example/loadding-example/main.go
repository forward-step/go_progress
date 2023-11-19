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
	var start int = 0
	var end int = 12
	var length int = end - start + 1
	var arr []*loading.Loading
	var isFirstRun bool = true
	for i := 0; i < length; i++ {
		arr = append(arr, loading.New(i))
	}
	for {
		var buf strings.Builder
		for i := start; i <= end; i++ {
			var item = arr[i]
			buf.WriteString(item.Next())
			buf.WriteByte('\n')
		}
		if !isFirstRun {
			tool.ClearLine(os.Stdout, length+1)
		}
		isFirstRun = false
		fmt.Print(buf.String())
		time.Sleep(time.Millisecond * 200)
	}
}
