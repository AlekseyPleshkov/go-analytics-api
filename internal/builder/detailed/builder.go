package compact

import (
	data "go-analytics/internal/data/detailed"
	domain "go-analytics/internal/domain/detailed"

	"gorm.io/gorm"
)

func Create(db *gorm.DB) (domain.Service, error) {
	store, err := data.NewStore(db)
	if err != nil {
		return nil, err
	}

	service := domain.NewService(store)

	return service, err
}
