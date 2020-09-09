package Entities

import "github.com/gin-gonic/gin"

type RoutesConfig struct {
	Engine *gin.Engine
}

func NewRoutesConfig() *RoutesConfig {
	return &RoutesConfig{
		Engine: gin.Default(),
	}
}
