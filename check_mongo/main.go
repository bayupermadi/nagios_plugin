package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"os"
)

type DBStats struct {
	DataSize    float64 `bson:"dataSize"`
	StorageSize float64 `bson:"storageSize"`
	IndexSize   float64 `bson:"indexSize"`
}

type ServerStats struct {
	Connection Connects `bson:"connections"`
}

type Connects struct {
	Current float64 `bson:"current"`
}

const (
	BytesPerGigabyte = 1073741824.0
)

var (
	config *Config
)

func main() {

	var status string
	var code int
	config = &Config{}
	parseCommandLine(config)

	// Set up connection
	session, err := mgo.Dial(config.DBHost)
	if err != nil {
		log.Fatal(err)
	}

	defer session.Close()

	// Get total database size
	if config.DBName != " " {
		db := session.DB(config.DBName)

		var dbStats DBStats
		if err := db.Run(bson.D{{"dbStats", 1}, {"scale", 1}}, &dbStats); err != nil {
			log.Fatal(err)
		}

		if config.Monitor == "DataSize" {
			if dbStats.DataSize/BytesPerGigabyte > config.Critical {
				status = "Critical"
				code = 2
			} else if dbStats.DataSize/BytesPerGigabyte > config.Warning {
				status = "Warning"
				code = 1
			} else {
				status = "OK"
				code = 0
			}

			fmt.Printf("%v - Data Size: %f GB;| DataSize=%d\n",
				status,
				dbStats.DataSize/BytesPerGigabyte,
				int(dbStats.DataSize))
			os.Exit(code)
		} else if config.Monitor == "StorageSize" {
			if dbStats.StorageSize/BytesPerGigabyte > config.Critical {
				status = "Critical"
				code = 2
			} else if dbStats.StorageSize/BytesPerGigabyte > config.Warning {
				status = "Warning"
				code = 1
			} else {
				status = "OK"
				code = 0
			}

			fmt.Printf("%v - Storage Size: %f GB;| StorageSize=%d\n",
				status,
				dbStats.StorageSize/BytesPerGigabyte,
				int(dbStats.StorageSize))
			os.Exit(code)
		} else if config.Monitor == "IndexSize" {
			if dbStats.IndexSize/BytesPerGigabyte > config.Critical {
				status = "Critical"
				code = 2
			} else if dbStats.IndexSize/BytesPerGigabyte > config.Warning {
				status = "Warning"
				code = 1
			} else {
				status = "OK"
				code = 0
			}

			fmt.Printf("%v - Index Size: %f GB;| IndexSize=%d\n",
				status,
				dbStats.IndexSize/BytesPerGigabyte,
				int(dbStats.IndexSize))
			os.Exit(code)
		} else {
			fmt.Print("Unknown Monitor")
			os.Exit(4)
		}
	} else {
		var serverStats ServerStats

		db := session.DB("admin")

		if err := db.Run(bson.D{{"serverStatus", 1}}, &serverStats); err != nil {
			panic(err)
			os.Exit(4)
		} else {
			if serverStats.Connection.Current > config.Critical {
				status = "Critical"
				code = 2
			} else if serverStats.Connection.Current > config.Warning {
				status = "Warning"
				code = 1
			} else {
				status = "OK"
				code = 0
			}

			fmt.Printf("%v - Total Cpnnection: %d;| Total Connection=%d\n",
				status,
				int(serverStats.Connection.Current),
				int(serverStats.Connection.Current))
		}
	}

}
