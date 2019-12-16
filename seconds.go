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
//
// Examples: next JSON:
//   {"time":1136239445}
// could be unmarshaled into next:
//   struct {
//     Time epoch.Seconds `json:"time"`
//   }
// and vice versa.
type Seconds struct {
	time.Time
}

// NewSeconds - returns Seconds
func NewSeconds(t time.Time) Seconds {
	return Seconds{Time: t}
}

// MarshalJSON - implements JSON marshaling interface
func (t Seconds) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Time.Unix())
}

// UnmarshalJSON - implements JSON unmarshaling interface
func (t *Seconds) UnmarshalJSON(data []byte) error {
	ts, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return errors.Wrap(err, "failed to parse int")
	}

	t.Time = time.Unix(ts, 0)

	return nil
}
