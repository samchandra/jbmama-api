package main

import "net/http"
import "github.com/gin-gonic/gin"


func main() {
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Listen and Server in 0.0.0.0:8081
	r.Run(":9081")
}
