package Routes

import "github.com/gin-gonic/gin"

type Route struct {
	PriceRoute *PriceRoute
}

func (route *Route) Handle(engine *gin.Engine) {
	route.PriceRoute.Handle(engine)
}

func NewRoute(priceRoute *PriceRoute) *Route {
	return &Route{
		PriceRoute: priceRoute,
	}
}
