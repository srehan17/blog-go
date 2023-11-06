package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/srehan17/go-crud/initializers"
	"github.com/srehan17/go-crud/models"
)

func PostsCreate(c *gin.Context) {

	// Get data from req body
	var body struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	c.Bind(&body)

	// Create a new post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post) // pass pointer of data to Create

	if result.Error != nil {
		c.Status(400)
		return
	}
	// Return the created post

	c.JSON(200, gin.H{
		"post": post,
	})
}
