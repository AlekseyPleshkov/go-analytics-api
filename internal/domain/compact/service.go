package compact

import (
	data "go-analytics/internal/data/compact"
)

type Service interface {
	Create(model *Compact) (*Compact, error)
}

type service struct {
	store data.Store
}

func NewService(store data.Store) Service {
	return &service{store: store}
}

func (s *service) Create(model *Compact) (*Compact, error) {
	scheme := mapModelToStoreScheme(model)
	if _, err := s.store.Create(scheme); err != nil {
		return nil, err
	}
	return model, nil
}
