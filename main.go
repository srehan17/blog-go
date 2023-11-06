package main

import (
	"github.com/gin-gonic/gin"
	"github.com/srehan17/go-crud/controllers"
	"github.com/srehan17/go-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)

	r.Run() // listen and serve on 0.0.0.0:8080
}
