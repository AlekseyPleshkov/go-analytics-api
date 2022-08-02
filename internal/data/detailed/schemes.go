package detailed

import (
	"database/sql"
	"time"

	"gorm.io/datatypes"
)

type Detailed struct {
	ID          uint      `gorm:"primarykey;unique;autoIncrement"`
	CreatedAt   time.Time `gorm:"not null"`
	UserID      string    `gorm:"not null"`
	Platform    string    `gorm:"not null"`
	Event       string    `gorm:"not null"`
	Category    sql.NullString
	Value       sql.NullString
	Additionals datatypes.JSON
}
