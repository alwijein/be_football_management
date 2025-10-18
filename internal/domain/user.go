package domain

import (
	"time"

	"gorm.io/gorm"
)

// User represents admin user
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Username  string         `gorm:"uniqueIndex;not null;size:50" json:"username"`
	Email     string         `gorm:"uniqueIndex;not null;size:100" json:"email"`
	Password  string         `gorm:"not null" json:"-"`
	FullName  string         `gorm:"size:100" json:"full_name"`
	PhotoURL  string         `gorm:"size:255" json:"photo_url"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName specifies table name
func (User) TableName() string {
	return "users"
}
