package progress

import (
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/term"
)

func termWidth() (width int, err error) {
	width, _, err = term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 0, err
	}
	return width, nil
}

func elapsedTimeByDuration(ms time.Duration) string {
	var res strings.Builder

	h := ms / time.Hour
	m := (ms % time.Hour) / time.Minute
	s := ms / time.Second
	if h > 0 {
		res.WriteString(strconv.Itoa(int(h)) + "h ")
	}
	if h > 0 || m > 0 {
		res.WriteString(strconv.Itoa(int(m)) + "m ")
	}
	res.WriteString(strconv.Itoa(int(s)) + "s")

	return res.String()
}
