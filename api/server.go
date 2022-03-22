package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/rtpa25/go_api_worflow/db/sqlc"
	"github.com/rtpa25/go_api_worflow/token"
	"github.com/rtpa25/go_api_worflow/utils"
)

//server set's up HTTP routing for our banking service
type Server struct {
	config     utils.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

//starts the server
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func (server *Server) setupRouter() {
	router := gin.Default()

	//unauthorized routes
	router.POST("/user", server.createUser)
	router.POST("/user/login", server.loginUser)

	//register middleware
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	//auth routes
	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccount)
	authRoutes.POST("/transfers", server.createTransfer)

	server.router = router
}

//newServer creates a new http server and set's up routing
func NewServer(config utils.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		store:      store,
		config:     config,
		tokenMaker: tokenMaker,
	}

	//way to add custom validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()
	return server, nil
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
