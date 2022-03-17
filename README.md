[![Stand With Ukraine](https://raw.githubusercontent.com/vshymanskyy/StandWithUkraine/main/banner-direct-single.svg)](https://vshymanskyy.github.io/StandWithUkraine)

# epoch

[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)
[![Godoc Reference][godoc-img]][godoc-url] [![build][build-img]][build-url] 
[![codecov][codecov-img]][codecov-url] [![goreportcard][goreportcard-img]][goreportcard-url]
[![Russian Warship Go Fuck Yourself](https://raw.githubusercontent.com/vshymanskyy/StandWithUkraine/main/badges/RussianWarship.svg)](https://vshymanskyy.github.io/StandWithUkraine)

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

## Additional terms of use for users from Russia and Belarus

By using the code provided in these repositories you agree with the following:
* Russia has [illegally annexed Crimea in 2014](https://en.wikipedia.org/wiki/Annexation_of_Crimea_by_the_Russian_Federation) and [brought the war in Donbas](https://en.wikipedia.org/wiki/War_in_Donbas) followed by [full-scale invasion of Ukraine in 2022](https://en.wikipedia.org/wiki/2022_Russian_invasion_of_Ukraine).
* Russia has brought sorrow and devastations to millions of Ukrainians, killed hundreds of innocent people, damaged thousands of buildings, and forced several million people to flee.
* [Putin khuylo!](https://en.wikipedia.org/wiki/Putin_khuylo!)

Glory to Ukraine! 🇺🇦

[godoc-img]: https://godoc.org/github.com/vtopc/epoch?status.svg
[godoc-url]: https://godoc.org/github.com/vtopc/epoch

[build-img]: https://github.com/vtopc/epoch/workflows/build/badge.svg
[build-url]: https://github.com/vtopc/epoch/actions?query=workflow%3A%22build%22

[codecov-img]: https://codecov.io/gh/vtopc/epoch/branch/master/graph/badge.svg
[codecov-url]: https://codecov.io/gh/vtopc/epoch

[goreportcard-img]: https://goreportcard.com/badge/github.com/vtopc/epoch
[goreportcard-url]: https://goreportcard.com/report/github.com/vtopc/epoch
