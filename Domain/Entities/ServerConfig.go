package Entities

type ServerConfig struct {
	Port string
}

func NewServerConfig() *ServerConfig {
	return &ServerConfig{
		Port: "8080",
	}
}
