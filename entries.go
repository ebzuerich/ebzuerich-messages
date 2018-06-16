package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Entry struct {
	Id int
	Title string
	Email string
	Message string

}

func createEntry(c *gin.Context) {

}

func indexEntries(c *gin.Context) {
	rows, err := conn.Query("SELECT * FROM entry")
	if err != nil {
		log.Fatalf("Error while quering entries: %q", err)
	}

	var entries []Entry
	var id int
	var title, email, message string
	for rows.Next() {
		err := rows.Scan(&id, &title, &email, &message)
		if err != nil {
			log.Fatalf("Error while reading entries: %q", err)
		}
		entries = append(entries, Entry{Id: id, Title: title, Email: email, Message: message})
	}

	c.JSON(200, entries)
}