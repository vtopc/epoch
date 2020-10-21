package epoch

import (
	"encoding/json"
	"fmt"
	"time"
)

// Milliseconds - same as epoch.Seconds, but for Epoch(Unix time) in milliseconds.
type Milliseconds struct {
	time.Time
}

// NewMilliseconds - returns Milliseconds
func NewMilliseconds(t time.Time) Milliseconds {
	return Milliseconds{Time: t}
}

// MarshalJSON - implements JSON marshaling interface
func (m Milliseconds) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Time.UnixNano() / nsPerMs)
}

// UnmarshalJSON - implements JSON unmarshaling interface
func (m *Milliseconds) UnmarshalJSON(data []byte) error {
	ms, err := parseInt64(string(data))
	if err != nil {
		return fmt.Errorf("failed to parse epoch.Milliseconds: %w", err)
	}

	m.Time = msToTime(ms)

	return nil
}

const (
	msPerS  = int64(time.Second / time.Millisecond)
	nsPerMs = int64(time.Millisecond)
)

func msToTime(ms int64) time.Time {
	s := ms / msPerS
	ns := (ms % msPerS) * nsPerMs

	return time.Unix(s, ns)
}
