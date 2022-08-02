package compact

import (
	domain "go-analytics/internal/domain/compact"
	"go-analytics/internal/routes/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRoutes(group *gin.RouterGroup) func(service domain.Service) {
	return func(service domain.Service) {
		group.POST("/compact", func(c *gin.Context) {
			request := request{}

			if err := c.BindJSON(&request); err != nil {
				c.JSON(
					http.StatusBadRequest,
					errors.NewError(errors.DecodeRequest, err),
				)
				return
			}

			model, err := mapRequestToModel(&request)

			if err != nil {
				c.JSON(
					http.StatusBadRequest,
					errors.NewError(errors.MapRequest, err),
				)
				return
			}

			if _, err := service.Create(model); err != nil {
				c.JSON(
					http.StatusInternalServerError,
					errors.NewError(errors.MapRequest, err),
				)
				return
			}

			c.JSON(http.StatusOK, response{})
		})
	}
}
