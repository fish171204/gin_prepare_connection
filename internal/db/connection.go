package db

// go get -u gorm.io/gorm
// go get -u gorm.io/driver/postgres

import "hoc-gin/internal/config"

var DB string

func InitDB() error {
	connStr := config.NewConfig().DNS()

	DB = connStr

	return nil
}
