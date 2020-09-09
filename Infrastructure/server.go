package server

import (
	Entity "../Domain/Entities"
)

type Server struct {
	Config *Entity.ServerConfig
	Routes *Routes
}

func (server *Server) Run() {
	engine := server.Routes.Handle()
	engine.Run()
}

func NewServer(serverConfig *Entity.ServerConfig, routes *Routes) *Server {
	return &Server{
		Config: serverConfig,
		Routes: routes,
	}
}
