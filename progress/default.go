package progress

import (
	"fmt"
	"os"
	"time"
)

func DefaultPrintRecord(r Record) string {
	switch r.Status {
	case STATUS_FAIL:
		return "fail"
	default:
		return fmt.Sprintf("%s %5.2f%%  %s  %d/%d", r.Rate, r.Percent, r.Time, r.Current, r.Total)
	}
}

func DefaultPrintAppendRecord(rs RecordSummary) string {
	return fmt.Sprintf("(%d/%d) %s %3.2f%%    %s    %d/%d", rs.SuccessNumber, rs.Length, rs.Rate, rs.Percent, rs.Time, rs.Current, rs.Total)
}

func defaultRateWidth(r RecordWithoutRate) int {
	return 30
}

func defaultPrintConfig() printConfig {
	return printConfig{
		// fill:  '█',
		// empty: ' ',
		fill:        '=',
		empty:       '-',
		head:        '>',
		left:        '[',
		right:       ']',
		writer:      os.Stdout,
		printRecord: DefaultPrintRecord,
		rateWidthFn: defaultRateWidth,
		refreshRate: time.Second,
	}
}
