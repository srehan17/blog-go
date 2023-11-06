package main

import (
	"github.com/srehan17/go-crud/initializers"
	"github.com/srehan17/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
