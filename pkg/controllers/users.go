package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUsers returns all users
func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GetUsers",
	})
}

func GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GetUser",
	})
}
