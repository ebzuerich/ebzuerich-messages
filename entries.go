package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Entry struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Email string `json:"email"`
	Message string `json:"message"`
	CreatedAt string `json:"created_at"`
}

func createEntry(c *gin.Context) {
	var newEntryId int
	err := conn.QueryRow(
		"INSERT INTO entry(title, email, message) VALUES ($1, $2, $3) returning id;",
		c.DefaultPostForm("title", "No title"),
		c.DefaultPostForm("email", "No email"),
		c.DefaultPostForm("message", "No message")).Scan(&newEntryId)

	if err != nil {
		log.Fatalf("Error while inserting entry: %q", err)
	}
}

func indexEntries(c *gin.Context) {
	rows, err := conn.Query("SELECT * FROM entry ORDER BY created_at DESC")
	if err != nil {
		log.Fatalf("Error while quering entries: %q", err)
	}

	var entries []Entry
	var id int
	var title, email, message, createdAt string
	for rows.Next() {
		err := rows.Scan(&id, &title, &email, &message, &createdAt)
		if err != nil {
			log.Fatalf("Error while reading entries: %q", err)
		}
		entries = append(entries, Entry{Id: id, Title: title, Email: email, Message: message, CreatedAt: createdAt})
	}

	c.JSON(200, entries)
}