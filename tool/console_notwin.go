//go:build !windows

package tool

import (
	"fmt"
	"os"
	"strings"
)

func ClearLine(file *os.File, line int) {
	const ESC = 27
	clear := fmt.Sprintf("%c[%dA%c[2K\r", ESC, 0, ESC)
	_, _ = fmt.Fprint(file, strings.Repeat(clear, line))
}
