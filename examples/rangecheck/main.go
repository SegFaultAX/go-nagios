package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/segfaultax/go-nagios"
)

var (
	warning  string
	critical string
	value    float64
	units    string
	noperf   bool
)

func init() {
	flag.StringVar(&warning, "warning", "", "warning range")
	flag.StringVar(&critical, "critical", "", "critical range")
	flag.Float64Var(&value, "value", 0.0, "the value to check")
	flag.StringVar(&units, "units", "kb", "unit of measure for value")
	flag.BoolVar(&noperf, "noperf", false, "disable perfdata")
}

func main() {
	flag.Parse()

	c, err := nagios.NewRangeCheckParse(warning, critical)
	if err != nil {
		fmt.Printf("error with check configuration: %s\n", err)
		os.Exit(3)
	}
	defer c.Done()

	c.CheckValue(value)

	if !noperf {
		c.AddPerfData(nagios.NewPerfData("value", value, units))
	}

	c.SetMessage("the current value is %f", value)
}
