package compact

import (
	"time"

	"gorm.io/datatypes"
)

type Compact struct {
	ID        uint           `gorm:"primarykey;unique;autoIncrement"`
	CreatedAt time.Time      `gorm:"not null"`
	UserID    string         `gorm:"not null"`
	Platform  string         `gorm:"not null"`
	Data      datatypes.JSON `gorm:"not null"`
}
