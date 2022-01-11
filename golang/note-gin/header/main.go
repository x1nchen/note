package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// Force log's color
	gin.ForceConsoleColor()

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		id := c.GetHeader("Sessionid")
		c.String(200, fmt.Sprintf("sessionid %s", id))
	})

	router.Run(":9000")
}
