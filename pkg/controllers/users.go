package controllers

import (
	"devbook-api/pkg/models"
	"devbook-api/pkg/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUsers is responsible for getting all users
func GetUsers(c *gin.Context) {

	// get all users
	users, err := repositories.UserRepository{}.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUser is responsible for getting a user by id
func GetUser(c *gin.Context) {

	// get the id from the request
	id := c.Param("id")

	// convert the id to integer
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// get the user by id
	user, err := repositories.UserRepository{}.FindByID(newid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

// CreateUser is responsible for creating a new user
func CreateUser(c *gin.Context) {

	// bind the request body to the user model
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// create a new user
	user, err := repositories.UserRepository{}.Create(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// UpdateUser is responsible for updating a user by id
func UpdateUser(c *gin.Context) {

	// get the id from the request
	id := c.Param("id")

	// convert the id to integer
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// bind the request body to the user model
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// update the user
	user, err = repositories.UserRepository{}.Update(newid, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUser is responsible for deleting a user by id
func DeleteUser(c *gin.Context) {

	// get the id from the request
	id := c.Param("id")

	// convert the id to integer
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// delete the user
	err = repositories.UserRepository{}.Delete(newid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
