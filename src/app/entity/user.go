package entity

import (
	// "fmt"
	"time"
	// "github.com/jinzhu/gorm"
)

// User ...
type User struct {
	ID        int    `json:"id" sql:"AUTO_INCREMENT"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
