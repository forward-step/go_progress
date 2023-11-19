package progress

import (
	"os"
	"time"
)

type Option func(*ProgressFactory)

func OptionRefreshRate(r time.Duration) Option {
	return func(pf *ProgressFactory) {
		pf.printConfig.refreshRate = r
	}
}

func OptionEnableDebug() Option {
	return func(pf *ProgressFactory) {
		pf.debug = true
	}
}

func OptionFailStrategy(fs FailStrategy) Option {
	return func(pf *ProgressFactory) {
		pf.failStrategy = fs
	}
}

func OptionOnFinish(fn FinshFn) Option {
	return func(pf *ProgressFactory) {
		pf.onFinsh = fn
	}
}

func OptionPrintRecord(fn PrintRecordFn) Option {
	return func(pf *ProgressFactory) {
		pf.printConfig.printRecord = fn
	}
}

func OptionPrintAppendRecord(fn PrintAppendRecordFn) Option {
	return func(pf *ProgressFactory) {
		pf.printConfig.printAppendRecordFn = fn
	}
}

func OptionPrintPrependRecord(fn PrintPrependRecordFn) Option {
	return func(pf *ProgressFactory) {
		pf.printConfig.printPrependRecordFn = fn
	}
}

func OptionRateWidth(w int) Option {
	return OptionRateWidthFn(func(RecordWithoutRate) int {
		return w
	})
}

func OptionRateWidthFn(w RateWidthFn) Option {
	return func(pf *ProgressFactory) {
		pf.printConfig.rateWidthFn = w
	}
}

func OptionHistoryLength(length uint) Option {
	return func(pf *ProgressFactory) {
		pf.historyLength = length
	}
}

func OptionPrintLeft(r rune) Option {
	return func(pf *ProgressFactory) {
		pf.printConfig.left = r
	}
}

func OptionPrintRight(r rune) Option {
	return func(pf *ProgressFactory) {
		pf.printConfig.right = r
	}
}

func OptionPrintFill(r rune) Option {
	return func(pf *ProgressFactory) {
		pf.printConfig.fill = r
	}
}

func OptionPrintEmpty(r rune) Option {
	return func(pf *ProgressFactory) {
		pf.printConfig.empty = r
	}
}

func OptionPrintHead(r rune) Option {
	return func(pf *ProgressFactory) {
		pf.printConfig.head = r
	}
}

func OptionPrintWriter(w *os.File) Option {
	return func(pf *ProgressFactory) {
		pf.printConfig.writer = w
	}
}
