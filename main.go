package main

import (
	"github.com/gin-gonic/gin"
	"github.com/srehan17/go-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world!",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
