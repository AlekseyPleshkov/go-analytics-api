package compact

import (
	data "go-analytics/internal/data/compact"
	"time"

	"gorm.io/datatypes"
)

func mapModelToStoreScheme(model *Compact) *data.Compact {
	return &data.Compact{
		CreatedAt: time.Now(),
		UserID:    model.UserID,
		Platform:  model.Platform,
		Data:      datatypes.JSON(model.Data),
	}
}
