package compact

import "gorm.io/gorm"

type Store interface {
	Create(scheme *Compact) (*Compact, error)
}

type store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) (Store, error) {
	if err := db.AutoMigrate(&Compact{}); err != nil {
		return nil, err
	}
	return &store{db: db}, nil
}

func (s *store) Create(scheme *Compact) (*Compact, error) {
	if result := s.db.Create(scheme); result.Error != nil {
		return nil, result.Error
	}
	return scheme, nil
}
