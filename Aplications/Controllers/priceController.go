package Controllers

import (
	Services "../../Domain/Services"
	"github.com/gin-gonic/gin"
)

type PriceController struct {
	priceService *Services.PriceService
}

func (priceController *PriceController) Get(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": priceController.priceService.Test(),
	})
}

func NewPriceController(_priceService *Services.PriceService) *PriceController {
	return &PriceController{
		priceService: _priceService,
	}
}
