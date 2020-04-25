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

	warn, err := nagios.ParseRange(warning)
	if err != nil {
		fmt.Printf("failed to configure check: %s", err)
		os.Exit(3)
	}

	crit, err := nagios.ParseRange(critical)
	if err != nil {
		fmt.Printf("failed to configure check: %s", err)
		os.Exit(3)
	}

	c := nagios.NewCheck()
	defer c.Done()

	if !noperf {
		c.AddPerfData(nagios.NewPerfData("value", value, units))
	}

	if crit.InRange(value) {
		c.Critical("we're in bad shape")
	} else if warn.InRange(value) {
		c.Warning("things are going sideways")
	} else {
		c.OK("everything looks good!")
	}
}
