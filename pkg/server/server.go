package server

import (
	"log"

	router "devbook-api/pkg/routes"

	"github.com/gin-gonic/gin"
)

// Server is a struct that contains the server port and the server router
type Server struct {
	port   string
	server *gin.Engine
}

// NewServer creates a new server
func NewServer(port string) *Server {
	return &Server{
		port:   port,
		server: gin.Default(),
	}
}

// Start starts the server
func (s *Server) Start() error {

	// Setup the routes
	r := router.SetupRouter(s.server)

	// Start the server
	log.Println("Server started on port", s.port)
	return r.Run("127.0.0.1:" + s.port)
}
