package detailed

import "gorm.io/gorm"

type Store interface {
	Create(scheme *Detailed) (*Detailed, error)
}

type store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) (Store, error) {
	if err := db.AutoMigrate(&Detailed{}); err != nil {
		return nil, err
	}
	return &store{db: db}, nil
}

func (s *store) Create(scheme *Detailed) (*Detailed, error) {
	if err := s.db.Create(scheme).Error; err != nil {
		return nil, err
	}
	return scheme, nil
}
