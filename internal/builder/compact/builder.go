package compact

import (
	data "go-analytics/internal/data/compact"
	domain "go-analytics/internal/domain/compact"

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
