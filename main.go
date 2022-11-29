package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello, world!",
	})
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/hello", Hello)
	err := router.Run()
	if err != nil {
		log.Fatalf("failed to start server: %s", err)
	}

}
