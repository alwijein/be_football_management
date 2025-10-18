package domain

import (
	"time"

	"gorm.io/gorm"
)

// Goal represents a goal scored in a match
type Goal struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	MatchID   uint           `gorm:"not null;index" json:"match_id"`
	PlayerID  uint           `gorm:"not null;index" json:"player_id"`
	Minute    int            `gorm:"not null" json:"minute"` // Menit terjadinya gol
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	
	// Relations
	Match  *Match  `gorm:"foreignKey:MatchID;constraint:OnDelete:CASCADE" json:"match,omitempty"`
	Player *Player `gorm:"foreignKey:PlayerID;constraint:OnDelete:CASCADE" json:"player,omitempty"`
}

// TableName specifies table name
func (Goal) TableName() string {
	return "goals"
}
