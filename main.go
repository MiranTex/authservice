package main

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	// "github.com/mirantex/gocontainer/src/model"
)

func main() {
	fmt.Println("great working")
	// model.Init()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("greet/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"message": "Hello " + toCaptalCase(name),
		})
	})

	r.Run(":8080")
}

func toCaptalCase(str string) string {
	return strings.Title(str)
}
