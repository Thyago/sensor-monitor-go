package controllers

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/thyago/sensor-monitor-go/config"
	"github.com/thyago/sensor-monitor-go/db"
)

type Server struct {
	router *gin.Engine
	db     *db.Database
	port   string
}

func NewServer(db *db.Database) *Server {
	// Create router
	r := gin.New()

	// Logger middleware
	r.Use(gin.Logger())

	// Recovery middleware: Recover from any panic and returns 500
	r.Use(gin.Recovery())

	return &Server{r, db, config.Config.ServerPort}
}

func (s *Server) Run() {
	// Initialize routes
	s.initRouter()
	s.router.Run(fmt.Sprintf(":%v", s.port))
}

func (s *Server) initRouter() {

	// TODO: Include auth

	// Configure cors
	s.router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))

	// Create API v1 routing endpoints
	v1 := s.router.Group("v1")
	{
		v1.POST("/sensors", s.CreateSensor)
		v1.GET("/sensors", s.ListSensor)
		v1.PUT("/sensors/:sensor-id", s.UpdateSensor)
		v1.GET("/sensors/:sensor-id", s.GetSensor)
		v1.DELETE("/sensors/:sensor-id", s.RemoveSensor)

		v1.GET("/sensors/:sensor-id/numeric-data", s.ListSensorDataNumeric)
		v1.DELETE("/sensors/:sensor-id/numeric-data", s.RemoveAllSensorDataNumeric)
	}
}
