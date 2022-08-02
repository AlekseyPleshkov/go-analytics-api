package detailed

import (
	"database/sql"
	data "go-analytics/internal/data/detailed"
	"time"

	"gorm.io/datatypes"
)

func mapModelToScheme(model *Detailed) *data.Detailed {
	category := sql.NullString{
		String: model.Category,
		Valid:  model.Category != "",
	}
	value := sql.NullString{
		String: model.Value,
		Valid:  model.Value != "",
	}

	return &data.Detailed{
		CreatedAt:   time.Now(),
		UserID:      model.UserID,
		Platform:    model.Platform,
		Event:       model.Event,
		Category:    category,
		Value:       value,
		Additionals: datatypes.JSON(model.Additionals),
	}
}
