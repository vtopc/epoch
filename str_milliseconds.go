package epoch

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// StrMilliseconds - same as epoch.Milliseconds, but for strings.
type StrMilliseconds struct {
	time.Time
}

// NewStrMilliseconds - returns Milliseconds
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
		return fmt.Errorf("failed to unmarshal epoch.StrMilliseconds: %w", err)
	}

	ms, err := parseInt64(v)
	if err != nil {
		return fmt.Errorf("failed to parse epoch.StrMilliseconds: %w", err)
	}

	m.Time = msToTime(ms)

	return nil
}
