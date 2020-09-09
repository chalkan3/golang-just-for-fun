package server

import (
	Route "../Aplications/Routes"
	Entity "../Domain/Entities"
	"github.com/gin-gonic/gin"
)

// Retorno minhas rotas para que eu possa usar elas prontas no meu server
type Routes struct {
	Routes *Entity.RoutesConfig
	Route  *Route.Route
}

func (routes *Routes) Handle() *gin.Engine {
	// Handle routes
	routes.Route.Handle(routes.Routes.Engine)
	return routes.Routes.Engine
}

func NewRoutes(routesConfig *Entity.RoutesConfig, route *Route.Route) *Routes {
	return &Routes{
		Routes: routesConfig,
		Route:  route,
	}
}
