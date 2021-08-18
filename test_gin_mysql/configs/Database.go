package configs

import (
	"sync"
)

type Database struct {
	Host     string
	Username string
	Password string
	Database string
	Port     int
}

var DatabaseConfigInstance *Database
var DatabaseConfigOnce sync.Once

func GetDatabaseConfig() *Database {
	DatabaseConfigOnce.Do(func() {
		DatabaseConfigInstance = &Database{
			Host:     "127.0.0.1",
			Username: "root",
			Password: "root",
			Database: "bayer_sdn",
			Port:     3306,
		}
	})
	return DatabaseConfigInstance
}
