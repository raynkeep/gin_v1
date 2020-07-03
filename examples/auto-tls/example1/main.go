package main

import (
	"log"

	"github.com/gin-gonic/autotls"
	"github.com/ryanker/gin_v1"
)

func main() {
	r := gin.Default()

	// Ping handler
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	log.Fatal(autotls.Run(r, "example1.com", "example2.com"))
}
