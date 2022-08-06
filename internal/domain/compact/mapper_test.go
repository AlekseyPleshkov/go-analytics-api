package compact

import (
	data "go-analytics/internal/data/compact"
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

	model := &Compact{
		UserID:   "user_id",
		Platform: "platform",
		Data:     "{\"source\": \"social\"}",
	}
	expected := &data.Compact{
		CreatedAt: timeMock(),
		UserID:    model.UserID,
		Platform:  model.Platform,
		Data:      datatypes.JSON(model.Data),
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
