package epoch_test

import (
	"encoding/json"
	"fmt"
	"github.com/vtopc/epoch"
)

type Request struct {
	Timestamp epoch.Seconds `json:"timestamp"`
}

func ExampleSeconds() {
	var v Request
	err := json.Unmarshal([]byte(`{"timestamp":1136239445}`), &v)
	if err != nil {
		panic(err)
	}

	// Also as epoch.Seconds inherits all time.Time's methods:
	fmt.Println(v.Timestamp.Year())
	fmt.Println(v.Timestamp.UTC().String())
	// Output:
	// 2006
	// 2006-01-02 22:04:05 +0000 UTC
}
