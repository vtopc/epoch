package epoch

import (
	"strconv"
	"time"
)

const (
	msPerS  = int64(time.Second / time.Millisecond)
	nsPerMs = int64(time.Millisecond)
)

func parseInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

func parseFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}
