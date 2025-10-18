package domain

import (
	"time"

	"gorm.io/gorm"
)

// PlayerPosition represents player position enum
type PlayerPosition string

const (
	PositionForward    PlayerPosition = "Penyerang"
	PositionMidfielder PlayerPosition = "Gelandang"
	PositionDefender   PlayerPosition = "Bertahan"
	PositionGoalkeeper PlayerPosition = "Penjaga Gawang"
)

// Player represents football player
type Player struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	TeamID       *uint          `gorm:"index" json:"team_id"`
	Name         string         `gorm:"not null;size:100" json:"name"`
	Height       int            `gorm:"not null" json:"height"` // in cm
	Weight       int            `gorm:"not null" json:"weight"` // in kg
	Position     PlayerPosition `gorm:"type:varchar(20);not null" json:"position"`
	JerseyNumber int            `gorm:"not null" json:"jersey_number"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	
	// Relations
	Team  *Team  `gorm:"foreignKey:TeamID;constraint:OnDelete:SET NULL" json:"team,omitempty"`
	Goals []Goal `gorm:"foreignKey:PlayerID;constraint:OnDelete:CASCADE" json:"-"`
}

// TableName specifies table name
func (Player) TableName() string {
	return "players"
}

// IsValidPosition checks if position is valid
func IsValidPosition(position string) bool {
	return position == string(PositionForward) ||
		position == string(PositionMidfielder) ||
		position == string(PositionDefender) ||
		position == string(PositionGoalkeeper)
}
