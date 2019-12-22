package epoch

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

// Seconds - seconds since the Epoch(Unix time).
// Inherits built-in time.Time type, thus has all it methods, but has custom serializer and
// deserializer(converts integer into built-in time.Time and vice versa).
type Seconds struct {
	time.Time
}

// NewSeconds - returns Seconds
func NewSeconds(t time.Time) Seconds {
	return Seconds{Time: t}
}

// MarshalJSON - implements JSON marshaling interface
func (s Seconds) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Time.Unix())
}

// UnmarshalJSON - implements JSON unmarshaling interface
func (s *Seconds) UnmarshalJSON(data []byte) error {
	ts, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return errors.Wrap(err, "failed to parse Seconds")
	}

	s.Time = time.Unix(ts, 0)

	return nil
}
