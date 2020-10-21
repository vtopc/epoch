package epoch

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// StrSeconds - seconds since the Epoch(Unix time) as string.
// Inherits built-in time.Time type, thus has all it methods, but has custom serializer and
// deserializer(converts integer into built-in time.Time and vice versa).
type StrSeconds struct {
	time.Time
}

// NewStrSeconds - returns StrSeconds
func NewStrSeconds(t time.Time) StrSeconds {
	return StrSeconds{Time: t}
}

// MarshalJSON - implements JSON marshaling interface
func (s StrSeconds) MarshalJSON() ([]byte, error) {
	return json.Marshal(strconv.FormatInt(s.Time.Unix(), 10))
}

// UnmarshalJSON - implements JSON unmarshaling interface
func (s *StrSeconds) UnmarshalJSON(data []byte) error {
	var v string

	err := json.Unmarshal(data, &v)
	if err != nil {
		return fmt.Errorf("failed to unmarshal epoch.StrSeconds: %w", err)
	}

	ts, err := parseInt64(v)
	if err != nil {
		return fmt.Errorf("failed to parse epoch.StrSeconds: %w", err)
	}

	s.Time = time.Unix(ts, 0)

	return nil
}
