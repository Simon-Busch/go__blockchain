package api

import (
	"github.com/go-kit/log"
	"github.com/labstack/echo/v4"
)


type ServerConfig struct {
	Logger        log.Logger
	ListenAddr				string
}

type Server struct {
	ServerConfig
}

func NewServer(cfg ServerConfig) *Server {
	return &Server{
		ServerConfig: cfg,
	}
}

func (s *Server) Start() error {
	e := echo.New()
	return e.Start(s.ListenAddr)
}
