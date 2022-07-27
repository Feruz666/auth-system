package api

import (
	"fmt"

	db "github.com/Feruz666/auth-system/db/sqlc"
	document "github.com/Feruz666/auth-system/pkg/document/handlers"
	sensors "github.com/Feruz666/auth-system/pkg/sensor/handlers"
	"github.com/Feruz666/auth-system/token"
	"github.com/Feruz666/auth-system/util"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSecretKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("/token/refresh", server.renewAccessToken)
	authRoutes.GET("/users", server.listUser)
	authRoutes.GET("/users/:id", server.getUser)

	// documents
	authRoutes.POST("/doc", document.Example)

	// Sensors route
	authRoutes.POST("/sensors", sensors.CreateSensor)
	authRoutes.GET("/sensors", sensors.GetSensors)
	authRoutes.GET("/sensors/charts", sensors.GetSensorsCharts)

	server.router = router
}

// Start runs
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
