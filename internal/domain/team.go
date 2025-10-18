package domain

import (
	"time"

	"gorm.io/gorm"
)

// Team represents football team
type Team struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	Name            string         `gorm:"not null;size:100" json:"name"`
	Logo            string         `gorm:"size:255" json:"logo"`
	EstablishedYear int            `gorm:"not null" json:"established_year"`
	Address         string         `gorm:"type:text" json:"address"`
	City            string         `gorm:"size:100" json:"city"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`

	// Relations
	Players     []Player `gorm:"foreignKey:TeamID;constraint:OnDelete:SET NULL" json:"players,omitempty"`
	HomeMatches []Match  `gorm:"foreignKey:HomeTeamID;constraint:OnDelete:CASCADE" json:"-"`
	AwayMatches []Match  `gorm:"foreignKey:AwayTeamID;constraint:OnDelete:CASCADE" json:"-"`
}

// TableName specifies table name
func (Team) TableName() string {
	return "teams"
}
