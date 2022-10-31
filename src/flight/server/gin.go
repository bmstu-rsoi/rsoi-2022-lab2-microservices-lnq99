package server

import (
	"app/flight/config"
	"app/flight/controller"
	"app/flight/service"
	"app/pkg/server"

	limits "github.com/gin-contrib/size"
	"github.com/gin-gonic/gin"
)

type GinServer struct {
	server.BaseServer
	engine   *gin.Engine
	handlers *controller.GinController
	config   *config.ServerConfig
}

func NewGinServer(service service.Service, cfg *config.ServerConfig) server.Server {
	// engine.SetMode(engine.ReleaseMode)
	engine := gin.New()

	// Middleware
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.Use(limits.RequestSizeLimiter(10))

	ctrl := controller.NewGinController(service)
	s := GinServer{
		engine:   engine,
		handlers: ctrl,
		config:   cfg,
	}
	s.SetupRouter()
	return &s
}

func (s *GinServer) SetupRouter() {
	r := s.engine
	r.GET("/api/v1/flights", s.handlers.ListFlights)
	//r.GET("/api/v1/flights/:id", s.handlers.GetFlight)
}

func (s *GinServer) Run() {
	cfg := server.DefaultConfig()
	cfg.Addr = s.config.Host + ":" + s.config.Port
	cfg.Handler = s.engine
	s.InitHttpServer(cfg)
	s.BaseServer.Run()
}
