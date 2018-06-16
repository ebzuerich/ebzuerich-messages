package main

import "database/sql"

var (
	conn *sql.DB
)

func main() {
	setupDatabase()
	setupRouting()
}