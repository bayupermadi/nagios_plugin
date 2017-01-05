package main

import (
	"fmt"
	"github.com/shirou/gopsutil/mem"
)

var (
	config *Config
)

func main() {
	v, _ := mem.VirtualMemory()

	var status string
	config = &Config{}
	parseCommandLine(config)

	if float64(v.UsedPercent) > config.Critical {
		status = "Critical"
	} else if float64(v.UsedPercent) > config.Warning {
		status = "Warning"
	} else {
		status = "OK"
	}

	// almost every return value is a struct
	//fmt.Printf("Total: %v, Free:%v, UsedPercent:%%%\n", v.Total, v.Free, v.UsedPercent)
	fmt.Printf("%v - Usage: %f%%;| Usage=%v\n", status, v.UsedPercent, v.Used)

	// convert to JSON. String() is also implemented
	//fmt.Println(v)
}
