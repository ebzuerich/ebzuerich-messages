package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setupRouting() {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/entry", indexEntries)
	router.Run()
}