package compact

import (
	domain "go-analytics/internal/domain/compact"
	"reflect"
	"testing"
)

func TestMapRequestToModel(t *testing.T) {
	// given
	request := &request{
		UserID:   "user_id",
		Platform: "platform",
		Data:     map[string]any{"screen": "main"},
	}
	expected := &domain.Compact{
		UserID:   request.UserID,
		Platform: request.Platform,
		Data:     "{\"screen\":\"main\"}",
	}

	// when
	result, _ := mapRequestToModel(request)

	// then
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("models should be equal, expects: %v, got: %v", expected, result)
	}
}
