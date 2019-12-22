package epoch

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

// Milliseconds - same as epoch.Seconds, but for Epoch(Unix time) in milliseconds.
type Milliseconds struct {
	time.Time
}

const (
	msPerS  = int64(time.Second / time.Millisecond)
	nsPerMs = int64(time.Millisecond)
)

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
	ms, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return errors.Wrap(err, "failed to parse Milliseconds")
	}

	s := ms / msPerS
	ns := (ms % msPerS) * nsPerMs

	m.Time = time.Unix(s, ns)

	return nil
}
