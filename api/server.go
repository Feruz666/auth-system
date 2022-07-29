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

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.GET("/users", server.listUser)
	authRoutes.GET("/users/:id", server.getUser)

	// documents
	authRoutes.POST("/doc", document.Example)

	// Sensors route
	authRoutes.POST("/sensors", sensors.CreateSensor)
	authRoutes.GET("/sensors", sensors.GetSensors)
	authRoutes.GET("/sensors/charts", sensors.GetSensorsCharts)

	// Maps route
	router.GET("/maps/workspaces", maps.GetWorkspaces)
	router.GET("/maps/workspaces/workspace", maps.GetWorkspace)
	router.POST("/maps/workspaces/workspace", maps.CreateWorkspace)
	router.DELETE("/maps/workspaces/workspace", maps.DeleteWorkspace)
	router.GET("/maps/layers", maps.GetLayers)
	router.GET("/maps/layers/layer", maps.GetLayer)
	router.DELETE("/maps/layers/layer", maps.DeleteLayers)
	router.GET("/maps/datastores", maps.GetDatastores)
	router.GET("/maps/datastores/datastore", maps.GetDatastore)
	router.GET("/maps/datastores/datastore/exists", maps.DatastoreExists)
	router.POST("/maps/datastores/datastore", maps.CreateDatastore)
	router.DELETE("/maps/datastores/datastore", maps.DeleteDatastore)
	router.GET("/maps/layergroups", maps.GetLayerGroups)
	router.GET("/maps/layergroups/layergroup", maps.GetLayerGroup)
	router.GET("/maps/featuretypes", maps.GetFeatureTypes)
	router.GET("/maps/featuretypes/featuretype", maps.GetFeatureType)
	router.DELETE("/maps/featuretypes/featuretype", maps.DeleteFeatureType)
	router.GET("/maps/wms", maps.GetWMS)
	router.GET("/maps/styles", maps.GetStyles)

	server.router = router
}

// Start runs
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
