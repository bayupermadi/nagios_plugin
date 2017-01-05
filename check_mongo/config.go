package main

import (
	"flag"
)

type Config struct {
	Critical float64
	Warning  float64
	DBName   string
	DBHost   string
	Monitor  string
}

func parseCommandLine(c *Config) {
	flag.Float64Var(&c.Critical, "lc", 0, "Critical Limit (int) ")
	flag.Float64Var(&c.Warning, "lw", 0, "Warning Limit (int)")
	flag.StringVar(&c.DBHost, "host", "localhost", "DB host")
	flag.StringVar(&c.DBName, "db", " ", "Database Name")
	flag.StringVar(&c.Monitor, "monitor", " ", "DataSize,StorageSize,IndexSize")

	flag.Parse()
}
