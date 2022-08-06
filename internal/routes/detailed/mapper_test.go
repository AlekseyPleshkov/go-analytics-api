package detailed

import (
	domain "go-analytics/internal/domain/detailed"
	"reflect"
	"testing"
)

func TestMapRequestToModel(t *testing.T) {
	// given
	request := &request{
		UserID:      "user_id",
		Platform:    "platform",
		Event:       "event",
		Category:    "category",
		Value:       "value",
		Additionals: map[string]any{"screen": "main"},
	}
	expected := &domain.Detailed{
		UserID:      request.UserID,
		Platform:    request.Platform,
		Event:       request.Event,
		Category:    request.Category,
		Value:       request.Value,
		Additionals: "{\"screen\":\"main\"}",
	}

	// when
	result, _ := mapRequestToModel(request)

	// then
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("models should be equal, expects: %v, got: %v", expected, result)
	}
}
