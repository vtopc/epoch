package epoch

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

// StrMilliseconds - same as epoch.Milliseconds, but for strings.
type StrMilliseconds struct {
	time.Time
}

// NewMilliseconds - returns Milliseconds
func NewStrMilliseconds(t time.Time) StrMilliseconds {
	return StrMilliseconds{Time: t}
}

// MarshalJSON - implements JSON marshaling interface
func (m StrMilliseconds) MarshalJSON() ([]byte, error) {
	ms := m.Time.UnixNano() / nsPerMs
	return json.Marshal(strconv.FormatInt(ms, 10))
}

// UnmarshalJSON - implements JSON unmarshaling interface
func (m *StrMilliseconds) UnmarshalJSON(data []byte) error {
	var v string

	err := json.Unmarshal(data, &v)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal epoch.StrMilliseconds")
	}

	ms, err := parseInt64(v)
	if err != nil {
		return errors.Wrap(err, "failed to parse epoch.StrMilliseconds")
	}

	m.Time = msToTime(ms)

	return nil
}
