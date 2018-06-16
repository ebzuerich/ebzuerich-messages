package main

import "github.com/gin-gonic/gin"

type Entry struct {
	Id int32
	Title string
	Email string
	Message string

}

func createEntry(entry Entry) {

}

func indexEntries(c *gin.Context) {

}