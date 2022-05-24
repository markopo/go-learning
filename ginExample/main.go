package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Person struct {
	Name string `form:"name"`
}

func main() {

	r := gin.Default()

	r.GET("/", func(context *gin.Context) {

		var person Person
		name := ""

		if context.ShouldBind(&person) == nil && person.Name != "" {
			name = "God day " + person.Name
		}

		context.JSON(200, gin.H{
			"message": "Hello from Gin! " + name + "!",
		})
	})

	r.GET("/ping-time", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"serverTime": time.Now().UTC(),
		})
	})

	r.Run(":6868")
}
