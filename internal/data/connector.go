package data

import (
	"fmt"
	"go-analytics/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnect(config config.Database) (*gorm.DB, error) {
	connection := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.Username,
		config.Password,
		config.Database,
	)

	return gorm.Open(postgres.Open(connection))
}
