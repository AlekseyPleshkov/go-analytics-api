package compact

import (
	data "go-analytics/internal/data/compact"
	"time"

	"gorm.io/datatypes"
)

var timeNow = func() time.Time {
	return time.Now()
}

func mapModelToScheme(model *Compact) *data.Compact {
	return &data.Compact{
		CreatedAt: timeNow(),
		UserID:    model.UserID,
		Platform:  model.Platform,
		Data:      datatypes.JSON(model.Data),
	}
}
