package main

import (
	"encoding/json"
	"fmt"

	"github.com/vtopc/epoch"
)

type Request struct {
	Timestamp epoch.Seconds `json:"timestamp"`
}

func main() {
	var v Request
	err := json.Unmarshal([]byte(`{"timestamp":1136239445}`), &v)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", v)
	// Output: {Timestamp:2006-01-03 00:04:05 +0200 EET}

	// Also as epoch.Seconds inherits all time.Time's methods:
	fmt.Println(v.Timestamp.Year())
	// Output: 2006
	fmt.Println(v.Timestamp.UTC().String())
	// Output: 2006-01-02 22:04:05 +0000 UTC
}
