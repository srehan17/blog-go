package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to the CRUD API"})
}
