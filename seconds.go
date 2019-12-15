package epoch

import (
	"encoding/json"
	"strconv"
	"time"
)

// Seconds - seconds since the Epoch(Unix time).
// Inherits built-in time.Time type, thus inherits all it methods, but has custom serializer and
// deserializer(converts integer into built-in time.Time and vice versa).
type Seconds struct {
	time.Time
}

// MarshalJSON - implements JSON marshaling interface
func (t Seconds) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Time.Unix())
}

// UnmarshalJSON - implements JSON unmarshaling interface
func (t *Seconds) UnmarshalJSON(data []byte) error {
	ts, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}

	t.Time = time.Unix(ts, 0)

	return nil
}
