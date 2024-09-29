package db

import (
	"database/sql"
	"fmt"
	"github.com/Vivirinter/scalable-api/internal/config"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(cfg *config.Config) error {
	dbInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", cfg.DbHost, cfg.DbUser, cfg.DbPassword, cfg.DbName)
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return err
	}

	DB = db
	return nil
}

func CloseDB() {
	DB.Close()
}
