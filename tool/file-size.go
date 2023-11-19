package tool

import (
	"fmt"
	"math"
)

// 计算文件大小
func FileSize(s float64) string {
	sizes := []string{"B", "kB", "MB", "GB", "TB", "PB", "EB"}
	base := 1024.0
	if s < 10 {
		return fmt.Sprintf("%2f%s", s, sizes[0])
	}
	e := math.Floor(logn(s, base))
	suffix := sizes[int(e)]
	val := math.Floor(s/math.Pow(base, e)*10+0.5) / 10
	f := "%.0f"
	if val < 10 {
		f = "%.1f"
	}

	return fmt.Sprintf("%s%s", fmt.Sprintf(f, val), suffix)
}

func logn(n, b float64) float64 {
	return math.Log(n) / math.Log(b)
}
