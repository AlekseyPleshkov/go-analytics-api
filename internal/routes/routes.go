package routes

import (
	"go-analytics/internal/config"
	domainCompact "go-analytics/internal/domain/compact"
	domainDetailed "go-analytics/internal/domain/detailed"

	routesCompact "go-analytics/internal/routes/compact"
	routesDetailed "go-analytics/internal/routes/detailed"

	"github.com/gin-gonic/gin"
)

func NewRoutes(
	config *config.Config,
	serviceCompact domainCompact.Service,
	serviceDetailed domainDetailed.Service,
) *gin.Engine {
	router := gin.Default()
	router.Use(tokenMiddlewere(config.Token))

	analytics := router.Group("/analytics")
	{
		v1 := analytics.Group("/v1")
		{
			routesCompact.NewRoutes(v1)(serviceCompact)
			routesDetailed.NewRoutes(v1)(serviceDetailed)
		}
	}

	return router
}
