package db

import (
	"fmt"
	"hoc-gin/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB string

func InitDB() error {
	connStr := config.NewConfig().DNS()

	conf, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return fmt.Errorf("error parsing DB config: %v", err)
	}

	return nil
}
