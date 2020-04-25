# go-nagios

This library is a small set of tools that make writing Nagios-style checks a bit
easier. The goal is for the primitives that this library exports to follow the
guidelines set forth in the [Nagios Plugin Development
Guidelines](https://nagios-plugins.org/doc/guidelines.html).

## Installation

`go get github.com/segfaultax/go-nagios`

## Usage

### Check

The `Check` type represents a very barebones Nagios check. A complete example of
how to use `Check` is in `examples/check`.

### RangeCheck

The `RangeCheck` type represents a class of Nagios checks that are specifically
checking ranges of numerical values. This logic is easy to implement yourself
with the normal `Check` type, but it's repetitive to do so. `RangeCheck` offers
a simple API for setting warning and critical ranges for your check. See a
complete example in `examples/rangecheck`.

### PerfData

The `PerfData` type represents Nagios performance data. Many monitoring systems
(such as Nagios and Sensu) will interpret this data and store it for later
querying. You can read about Nagios performance data
[here](http://nagios-plugins.org/doc/guidelines.html#AEN200). Both the `Check`
and `RangeCheck` types supporting adding arbitrary performance data to the check
results. See the `examples/` directory on how to do that.

### Range

The `Range` type represents a Nagios threshold range.

Ref: [Nagios Documentation](https://nagios-plugins.org/doc/guidelines.html#THRESHOLDFORMAT)


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
