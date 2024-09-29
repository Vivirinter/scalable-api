package migrations

import (
	"github.com/Vivirinter/scalable-api/internal/db"
	"log"
)

const (
	Migration = `CREATE TABLE IF NOT EXISTS scalable (
id serial PRIMARY KEY,
author text NOT NULL,
content text NOT NULL,
created_at timestamp with time zone DEFAULT current_timestamp
`
)

func RunMigrations() {
	_, err := db.DB.Query(Migration)
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
}
