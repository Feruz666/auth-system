package api

import (
	"fmt"
	"time"

	db "github.com/Feruz666/auth-system/db/sqlc"
	document "github.com/Feruz666/auth-system/pkg/document/handlers"
	maps "github.com/Feruz666/auth-system/pkg/maps/handlers"
	sensors "github.com/Feruz666/auth-system/pkg/sensor/handlers"
	"github.com/Feruz666/auth-system/token"
	"github.com/Feruz666/auth-system/util"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
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

	// Apply the middleware to the router (works with groups too)
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	router.POST("/token/refresh", server.renewAccessToken)
	router.POST("/token/verify", server.getUserByToken)

	// Sensors route
	router.POST("/sensors", sensors.CreateSensor)
	router.GET("/sensors", sensors.GetSensors)
	router.GET("/sensors/charts", sensors.GetSensorsCharts)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.GET("/users", server.listUser)
	authRoutes.GET("/users/:id", server.getUser)

	// documents
	authRoutes.POST("/doc", document.Example)

	// Maps route
	router.GET("/maps/workspaces", maps.MirrorGET)
	router.GET("/maps/workspaces/workspace", maps.MirrorGET)
	router.POST("/maps/workspaces/workspace", maps.CreateWorkspace)
	router.DELETE("/maps/workspaces/workspace", maps.DeleteWorkspace)
	router.GET("/maps/layers", maps.MirrorGET)
	router.GET("/maps/layers/layer", maps.MirrorGET)
	router.DELETE("/maps/layers/layer", maps.DeleteLayers)
	router.GET("/maps/datastore", maps.MirrorGET)
	router.GET("/maps/datastore/detail", maps.MirrorGET)
	router.GET("/maps/datastore/datastore/exists", maps.MirrorGET)
	router.POST("/maps/datastore/datastore", maps.CreateDatastore)
	router.DELETE("/maps/datastores/datastore", maps.DeleteDatastore)
	router.GET("/maps/layergroups", maps.MirrorGET)
	router.GET("/maps/layergroups/layergroup", maps.MirrorGET)
	router.GET("/maps/featuretypes", maps.MirrorGET)
	router.GET("/maps/featuretypes/featuretype", maps.MirrorGET)
	router.DELETE("/maps/featuretypes/featuretype", maps.DeleteFeatureType)
	router.GET("/maps/wms", maps.MirrorWMS)
	router.GET("/maps/styles", maps.MirrorGET)
	router.GET("/maps/coverage/dir", maps.MirrorGET)
	router.GET("/maps/coverage/info", maps.MirrorGET)
	router.GET("/maps/coverage/layers", maps.MirrorGET)
	router.POST("/maps/coverage/coveragestore", maps.CreateCoverageStore)
	router.POST("/maps/coverage/publish", maps.PublishCoverageLayer)

	server.router = router
}

// Start runs
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
