package progress

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/forward-step/go_progress/tool"
)

type PrintRecordFn func(Record) string
type PrintAppendRecordFn func(RecordSummary) string
type PrintPrependRecordFn func(RecordSummary) string
type RateWidthFn func(RecordWithoutRate) int

type printConfig struct {
	fill                 rune          // 填充内容
	head                 rune          // 头部内容
	empty                rune          // 空格内容
	left                 rune          // 开始符
	right                rune          // 结束符
	rateWidthFn          RateWidthFn   // 计算rate的宽度
	writer               *os.File      // 输出 ; os.Stdout
	refreshRate          time.Duration // 刷新频率
	printRecord          PrintRecordFn
	printAppendRecordFn  PrintAppendRecordFn
	printPrependRecordFn PrintPrependRecordFn
}

func (pf *ProgressFactory) refreshConsole() {
	pf.RLock()
	defer pf.RUnlock()

	if pf.updateCh == nil {
		return
	}

	var buf strings.Builder
	var endNumber uint8
	var arrLength = len(pf.arr)
	var recordSummary = RecordSummary{}
	var maxTime time.Duration
	var loadingIndex int

	// single
	for i, p := range pf.arr {
		record := p.getRecord(i)
		if p.isEnd() {
			record.LoadingIndex = -1
		} else {
			record.LoadingIndex = loadingIndex
			loadingIndex++
		}
		recordString := pf.printConfig.printRecord(record)
		if recordString != "" {
			buf.WriteString(recordString + "\n")
		}
		if p.isEnd() {
			endNumber++
		}

		recordSummary.Current += record.Current
		recordSummary.Total += record.Total
		maxTime = max(p.elapsedTime, maxTime)
		switch p.status {
		case STATUS_FAIL:
			recordSummary.FailNumber++
		case STATUS_SUCCESS:
			recordSummary.SuccessNumber++
		default:
		}
	}

	isGameOver := endNumber >= uint8(arrLength)

	recordSummary.Length = uint8(arrLength)

	switch {
	case isGameOver:
		if recordSummary.FailNumber > 0 {
			recordSummary.Status = STATUS_FAIL
		} else {
			recordSummary.Status = STATUS_SUCCESS
		}
	default:
		recordSummary.Status = STATUS_LOADING
	}

	if pf.printConfig.printAppendRecordFn != nil || pf.printConfig.printPrependRecordFn != nil {
		recordSummary.ElapsedTime = maxTime
		recordSummary.Time = elapsedTimeByDuration(maxTime)
		recordSummary.Percent = float64(recordSummary.Current) / float64(recordSummary.Total) * 100
		recordSummary.Rate = recordSummary.print(pf.printConfig)
	}

	// prepend
	if pf.printConfig.printPrependRecordFn != nil {
		var newBuf = strings.Builder{}
		newBuf.WriteString(pf.printConfig.printPrependRecordFn(recordSummary) + "\n")
		newBuf.WriteString(buf.String())
		buf = newBuf
	}

	// append
	if pf.printConfig.printAppendRecordFn != nil {
		buf.WriteString(pf.printConfig.printAppendRecordFn(recordSummary) + "\n")
	}

	// print
	if !pf.debug && len(pf.arr) != 0 {
		tool.ClearLine(pf.printConfig.writer, pf.line)
		_, _ = fmt.Fprint(pf.printConfig.writer, buf.String())
	}

	// all end
	if isGameOver {
		pf.finish(Result{
			Length:        recordSummary.Length,
			FailNumber:    recordSummary.FailNumber,
			SuccessNumber: recordSummary.SuccessNumber,
		})
		pf.Done <- Result{
			Length:        recordSummary.Length,
			FailNumber:    recordSummary.FailNumber,
			SuccessNumber: recordSummary.SuccessNumber,
		}
		return
	}

	// line
	var line int
	var currentLineLength int
	for _, r := range buf.String() {
		if r == '\n' {
			line++
			currentLineLength = 0
		} else if r == '\b' {
			currentLineLength = min(currentLineLength-1, 0)
		} else if r == '\t' {
			// 长度无法确定，为了安全打印进度条，从而多删除一行
			// currentLineLength += 4
			line++
			currentLineLength = 0
		} else if r == '\r' {
			currentLineLength = 0
		} else {
			currentLineLength++
		}
		if currentLineLength > TermWidth {
			line++
			currentLineLength = TermWidth - currentLineLength
		}
	}
	pf.line = line + 1
}

func (r *RecordWithoutRate) print(cfg printConfig) string {
	var rate strings.Builder
	var width = cfg.rateWidthFn(*r)
	{
		realWidth := width - 2
		fillWidth := int(r.Percent / 100 * float64(realWidth))
		emptyWidth := realWidth - fillWidth
		headWidth := 0
		if cfg.head != 0 && fillWidth > 0 && fillWidth < realWidth {
			fillWidth--
			headWidth++
		}
		// left
		rate.WriteRune(cfg.left)
		// fill
		for i := 1; i <= fillWidth; i++ {
			rate.WriteRune(cfg.fill)
		}
		// head
		for i := 1; i <= headWidth; i++ {
			rate.WriteRune(cfg.head)
		}
		// empty
		for i := 1; i <= emptyWidth; i++ {
			rate.WriteRune(cfg.empty)
		}
		// right
		rate.WriteRune(cfg.right)
	}

	return rate.String()
}
