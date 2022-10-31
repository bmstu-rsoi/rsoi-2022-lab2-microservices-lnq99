package controller

import (
	"app/flight/service"
	errors "app/pkg/error"
	"net/http"

	"app/pkg/util"

	"github.com/gin-gonic/gin"
)

// GinController converts gin contexts to parameters.
type GinController struct {
	service service.Service
}

func NewGinController(service service.Service) *GinController {
	return &GinController{service}
}

func (c *GinController) ListFlights(ctx *gin.Context) {
	page := util.ToInt(ctx.Query("page"))
	size := util.ToInt(ctx.Query("size"))

	if page <= 0 || size <= 0 {
		ctx.JSON(http.StatusBadRequest, errors.ErrorResponse{"Invalid params"})
	}

	r := c.service.ListFlights(ctx, int32(page), int32(size))
	ctx.JSON(http.StatusOK, r)
}
