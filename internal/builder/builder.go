package builder

import (
	builderCompact "go-analytics/internal/builder/compact"
	builderDetailed "go-analytics/internal/builder/detailed"
	"go-analytics/internal/config"
	"go-analytics/internal/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Create(config *config.Config, db *gorm.DB) (*gin.Engine, error) {
	serviceCompact, err := builderCompact.Create(db)

	if err != nil {
		return nil, err
	}

	serviceDetailed, err := builderDetailed.Create(db)

	if err != nil {
		return nil, err
	}

	routes := routes.NewRoutes(config, serviceCompact, serviceDetailed)

	return routes, nil
}
