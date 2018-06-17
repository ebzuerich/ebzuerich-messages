package main

import (
	"database/sql"
	"os"
	"log"
	_ "github.com/lib/pq"
)

func setupDatabase() {
	var err error

	conn, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	if _, err := conn.Exec(`CREATE TABLE IF NOT EXISTS entry (
id SERIAL PRIMARY KEY,
title varchar(100),
email varchar(100),
message text,
created_at timestamptz NOT NULL DEFAULT NOW()
);`); err != nil {
		log.Fatalf("Error creating database table: %q", err)
		return
	}
}