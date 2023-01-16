package router

import (
	"net/http"

	"devbook-api/pkg/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter is responsible for setting up the routes
func SetupRouter(router *gin.Engine) *gin.Engine {

	main := router.Group("/api/v1")

	main.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	users := main.Group("/users")
	{
		users.GET("/", controllers.GetUsers)
		users.GET("/:id", controllers.GetUser)
	}

	return router
}
