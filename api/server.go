package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/rtpa25/go_api_worflow/db/sqlc"
	"github.com/rtpa25/go_api_worflow/token"
)

//server set's up HTTP routing for our banking service
type Server struct {
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

//starts the server
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

//newServer creates a new http server and set's up routing
func NewServer(store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker("")
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{store: store}
	router := gin.Default()

	//way to add custom validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	//routes
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	router.POST("/transfers", server.createTransfer)

	router.POST("/user", server.createUser)

	server.router = router
	return server
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
