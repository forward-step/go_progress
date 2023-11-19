package main

import (
	"fmt"
	"time"

	"github.com/forward-step/go_progress/progress"
)

type ProgressData struct {
	title    string
	subtitle string
}

func main() {
	factory := progress.New(
		progress.OptionPrintRecord(func(r progress.Record) string {
			if data, ok := r.Data.(ProgressData); ok {
				switch r.Status {
				case progress.STATUS_FAIL:
					return "下载失败"
				default:
					rate := fmt.Sprintf("%.2f it/s", r.CurrentRate())
					return fmt.Sprintf("%s\n %s\n%s %5.2f%%    %s    %s    %d/%d", data.title, data.subtitle, r.Rate, r.Percent, rate, r.Time, r.Current, r.Total)
				}
			}
			return progress.DefaultPrintRecord(r)
		}),
	)
	p := factory.Add(100).Mount(ProgressData{
		title:    "title",
		subtitle: "second title",
	})

	for i := 0; i < 100; i++ {
		p.Add(1)
		time.Sleep(time.Millisecond * 100)
	}

	<-factory.Done
}
