package detailed

import (
	"errors"
	data "go-analytics/internal/data/detailed"
	"testing"
	"time"
)

func TestCreate_ShouldSuccess(t *testing.T) {
	// given
	storeMock := storeMock{}
	timeMock := func() time.Time {
		return time.Date(2022, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
	}

	service := NewService(&storeMock)
	model := &Detailed{}

	timeNow = timeMock

	// when
	_, err := service.Create(model)

	// then
	if err != nil {
		t.Errorf("error should be nil, got %v", err)
	}

	if storeMock.createReceivedScheme == nil {
		t.Errorf("received scheme shouldn't be nil")
	}
}

func TestCreate_ShouldFail(t *testing.T) {
	// given
	storeMock := storeMock{}

	service := NewService(&storeMock)
	model := &Detailed{}
	expected := errors.New("example error")

	storeMock.createErrorStub = expected

	// when
	result, err := service.Create(model)

	// then
	if err != expected {
		t.Errorf("error shouldn't be nil")
	}

	if result != nil {
		t.Errorf("result should be nil")
	}
}

// Mocks

type storeMock struct {
	createWasCalled      int
	createReceivedScheme *data.Detailed
	createErrorStub      error
}

func (s *storeMock) Create(scheme *data.Detailed) (*data.Detailed, error) {
	s.createWasCalled += 1
	s.createReceivedScheme = scheme
	return nil, s.createErrorStub
}
