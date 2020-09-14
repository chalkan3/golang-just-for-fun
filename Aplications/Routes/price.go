package Routes

import (
	Helpers "../../Domain/Helpers"
	Controllers "../Controllers"
	"github.com/gin-gonic/gin"
)

type PriceRoute struct {
	baseuri          string
	priceController  *Controllers.PriceController
	handleController *Helpers.HandleController
}

func (priceRoute *PriceRoute) Handle(engine *gin.Engine) {
	priceRoute.get(engine)
}

func (priceRoute *PriceRoute) get(engine *gin.Engine) {
	engine.GET(priceRoute.baseuri, priceRoute.handleController.Call(priceRoute.priceController.Get))
}

func NewPriceRoute(priceController *Controllers.PriceController, handleController *Helpers.HandleController) *PriceRoute {
	return &PriceRoute{
		baseuri:          "/price",
		priceController:  priceController,
		handleController: handleController,
	}
}
