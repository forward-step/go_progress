package main

import (
	"io"
	"net/http"
	"os"

	"github.com/forward-step/go_progress/progress"
)

func main() {
	req, err := http.NewRequest("GET", "https://dl.google.com/go/go1.14.2.src.tar.gz", nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	file, err := os.OpenFile("go1.14.2.src.tar.gz", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	f := progress.New()
	p := f.Add(resp.ContentLength)
	defer p.Close()
	io.Copy(io.MultiWriter(file, p), resp.Body)
}
