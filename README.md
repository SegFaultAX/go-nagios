# go-nagios

This library is a small set of tools that make writing Nagios-style checks a bit
easier. The goal is for the primitives that this library exports to follow the
guidelines set forth in the [Nagios Plugin Development
Guidelines](https://nagios-plugins.org/doc/guidelines.html).

## Installation

`go get github.com/segfaultax/go-nagios`

## Usage

### Range

Ref: [Nagios Documentation](https://nagios-plugins.org/doc/guidelines.html)

The Range type represents a Nagios threshold range.


```go
import "github.com/segfaultax/go-nagios"

r1 := nagios.NewRange(0, 100, false) // exclusive range outside of 0 - 100
r1.InRange(50) // false
r1.InRange(250) // true

r2 := nagios.NewRange(0, 100, true) // inclusive range between 0 - 100
r2.InRange(50) // true
r2.InRange(250) // false
```

Range also supports parsing a Nagios-style threshold string.

```go
import "github.com/segfaultax/go-nagios"

r1 := nagios.ParseRange("100") // exclusive range outside of 0 - 100
r1.InRange(50) // false
r1.InRange(250) // true

r2 := nagios.ParseRange("@100") // inclusive range between 0 - 100
r1.InRange(50) // false
r1.InRange(250) // true
```

See the Nagios documentation for a full description of the range format
specification.

## Contributors

* Michael-Keith Bernard (@segfaultax)

## License

This project is released under the MIT license. See LICENSE for details.
