package epoch

import (
	"encoding/json"
	"fmt"
	"math"
	"time"
)

// FloatMS - integer part of timestamp represents seconds and fractional - milliseconds since
// the Epoch(Unix time), e.g.
//   1136239445.999
//
// Inherits built-in time.Time type, thus has all it methods, but has custom serializer and
// deserializer(converts float timestamp into built-in time.Time and vice versa).
type FloatMS struct {
	time.Time
}

// NewFloatMS - returns FloatMS
func NewFloatMS(t time.Time) FloatMS {
	return FloatMS{Time: t}
}

// MarshalJSON - implements JSON marshaling interface
func (s FloatMS) MarshalJSON() ([]byte, error) {
	milli := s.Time.UnixMilli()
	f := float64(milli) / float64(msPerS)

	return json.Marshal(f)
}

func (s *FloatMS) UnmarshalJSON(data []byte) error {
	f, err := parseFloat64(string(data))
	if err != nil {
		return fmt.Errorf("failed to parse epoch.FloatMS: %w", err)
	}

	i, frac := math.Modf(f)
	s.Time = time.Unix(int64(i), int64(frac*1_000)*nsPerMs)

	return nil
}
