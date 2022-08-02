package detailed

import (
	data "go-analytics/internal/data/detailed"
)

type Service interface {
	Create(model *Detailed) (*Detailed, error)
}

type service struct {
	store data.Store
}

func NewService(store data.Store) Service {
	return &service{store: store}
}

func (s *service) Create(model *Detailed) (*Detailed, error) {
	scheme := mapModelToScheme(model)

	if _, err := s.store.Create(scheme); err != nil {
		return nil, err
	}

	return model, nil
}
