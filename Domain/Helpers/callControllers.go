package Helpers

import (
	"github.com/gin-gonic/gin"
)

type HandleController struct {
}

func (handleController *HandleController) Call(controller func(c *gin.Context)) func(c *gin.Context) {
	return func(c *gin.Context) { controller(c) }
}

func NewHandleController() *HandleController {
	return &HandleController{}
}
