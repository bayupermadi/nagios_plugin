package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"syscall"
	"time"
)

var (
	config *Config
)

func main() {

	var status string
	config = &Config{}
	parseCommandLine(config)

	cpu_res, e := cpu.Percent(time.Duration(1)*time.Second, false)
	if e == nil {
		if cpu_res[0] > config.Critical {
			status = "Critical"
		} else if cpu_res[0] > config.Warning {
			status = "Warning"
		} else {
			status = "OK"
		}
	} else {
		fmt.Printf("Error getting cpu usage: %s\n", e.Error())
		syscall.Exit(1)
	}

	// almost every return value is a struct
	//fmt.Printf("Total: %v, Free:%v, UsedPercent:%%%\n", v.Total, v.Free, v.UsedPercent)
	fmt.Printf("%v - Usage: %v%%;| Usage=%v\n", status, cpu_res[0], cpu_res[0])

	// convert to JSON. String() is also implemented
	//fmt.Println(v)

}
