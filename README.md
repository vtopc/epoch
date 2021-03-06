# epoch

[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)
[![Godoc Reference][godoc-img]][godoc-url] [![build][build-img]][build-url] 
[![codecov][codecov-img]][codecov-url] [![goreportcard][goreportcard-img]][goreportcard-url]

Contains primitives for marshaling/unmarshaling Unix timestamp/epoch to/from built-in time.Time type in JSON.

## Seconds
Seconds since the Epoch(Unix time), e.g.:
```json
  {"timestamp":1136239445}
```
Inherits built-in time.Time type, thus has all it methods, but has custom serializer and
deserializer(converts integer into built-in time.Time and vice versa).

#### Usage Example

```go
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

	// Also as epoch.Seconds inherits all time.Time's methods one can do next:
	fmt.Println(v.Timestamp.Year())
	// Output: 2006
	fmt.Println(v.Timestamp.UTC().String())
	// Output: 2006-01-02 22:04:05 +0000 UTC
}
```

## Milliseconds
Same as epoch.Seconds, but for Epoch(Unix time) in milliseconds, e.g.:
```json
  {"timestamp":1136239445999}
```

## StrSeconds
Same as epoch.Seconds, but for strings, e.g.:
```json
  {"timestamp":"1136239445"}
```

## StrMilliseconds
Same as epoch.Milliseconds, but for strings, e.g.:
```json
  {"timestamp":"1136239445999"}
```

[godoc-img]: https://godoc.org/github.com/vtopc/epoch?status.svg
[godoc-url]: https://godoc.org/github.com/vtopc/epoch

[build-img]: https://github.com/vtopc/epoch/workflows/build/badge.svg
[build-url]: https://github.com/vtopc/epoch/actions?query=workflow%3A%22build%22

[codecov-img]: https://codecov.io/gh/vtopc/epoch/branch/master/graph/badge.svg
[codecov-url]: https://codecov.io/gh/vtopc/epoch

[goreportcard-img]: https://goreportcard.com/badge/github.com/vtopc/epoch
[goreportcard-url]: https://goreportcard.com/report/github.com/vtopc/epoch
