package Routes

import "github.com/gin-gonic/gin"

type PriceRoute struct {
	baseuri string
}

func (priceRoute *PriceRoute) Handle(engine *gin.Engine) {
	priceRoute.get(engine)
}

func (priceRoute *PriceRoute) get(engine *gin.Engine) {
	engine.GET(priceRoute.baseuri, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

func NewPriceRoute() *PriceRoute {
	return &PriceRoute{
		baseuri: "/price",
	}
}
