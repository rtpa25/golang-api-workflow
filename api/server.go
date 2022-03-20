package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/rtpa25/go_api_worflow/db/sqlc"
)

//server set's up HTTP routing for our banking service
type Server struct {
	store  db.Store
	router *gin.Engine
}

//starts the server
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

//newServer creates a new http server and set's up routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	//routes
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	server.router = router
	return server
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
