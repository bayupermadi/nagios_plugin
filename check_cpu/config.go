package main

import (
	"flag"
)

type Config struct {
	Critical float64
	Warning  float64
}

func parseCommandLine(c *Config) {
	flag.Float64Var(&c.Critical, "lc", 85, "Critical Limit (%)")
	flag.Float64Var(&c.Warning, "lw", 75, "Warning Limit (%)")

	flag.Parse()
}
