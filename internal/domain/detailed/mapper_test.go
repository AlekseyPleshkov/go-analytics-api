package detailed

import (
	"database/sql"
	data "go-analytics/internal/data/detailed"
	"reflect"
	"testing"
	"time"

	"gorm.io/datatypes"
)

func TestMapModelToStoreScheme(t *testing.T) {
	// given
	timeMock := func() time.Time {
		return time.Date(2022, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
	}

	model := &Detailed{
		UserID:      "user_id",
		Platform:    "platform",
		Event:       "event",
		Category:    "category",
		Value:       "value",
		Additionals: "{\"screen\": \"home\"}",
	}
	expected := &data.Detailed{
		CreatedAt:   timeMock(),
		UserID:      model.UserID,
		Platform:    model.Platform,
		Event:       model.Event,
		Category:    sql.NullString{String: model.Category, Valid: true},
		Value:       sql.NullString{String: model.Value, Valid: true},
		Additionals: datatypes.JSON(model.Additionals),
	}

	timeNow = timeMock

	// when
	result := mapModelToScheme(model)

	// then
	if !reflect.DeepEqual(result, expected) {
		t.Errorf(
			"models should be equal, expects: %+v, got: %+v",
			expected,
			result,
		)
	}
}
