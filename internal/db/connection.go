package db

// go get -u gorm.io/gorm
// go get -u gorm.io/driver/postgres

import (
	"fmt"
	"hoc-gin/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() error {
	connStr := config.NewConfig().DNS()

	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	var err error
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN: connStr,
	}), config)

	if err != nil {
		return fmt.Errorf("error opening DB connection: %w", err)
	}

	return nil
}
