package domain

import (
	"time"

	"gorm.io/gorm"
)

// MatchStatus represents match status
type MatchStatus string

const (
	StatusScheduled MatchStatus = "Scheduled"
	StatusCompleted MatchStatus = "Completed"
	StatusCancelled MatchStatus = "Cancelled"
)

// Match represents football match
type Match struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	HomeTeamID uint           `gorm:"not null;index" json:"home_team_id"`
	AwayTeamID uint           `gorm:"not null;index" json:"away_team_id"`
	MatchDate  time.Time      `gorm:"not null;index" json:"match_date"`
	MatchTime  string         `gorm:"not null;size:5" json:"match_time"`
	HomeScore  *int           `json:"home_score"`
	AwayScore  *int           `json:"away_score"`
	Status     MatchStatus    `gorm:"type:varchar(20);default:'Scheduled'" json:"status"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`

	// Relations
	HomeTeam *Team  `gorm:"foreignKey:HomeTeamID;constraint:OnDelete:CASCADE" json:"home_team,omitempty"`
	AwayTeam *Team  `gorm:"foreignKey:AwayTeamID;constraint:OnDelete:CASCADE" json:"away_team,omitempty"`
	Goals    []Goal `gorm:"foreignKey:MatchID;constraint:OnDelete:CASCADE" json:"goals,omitempty"`
}

// TableName specifies table name
func (Match) TableName() string {
	return "matches"
}

// GetMatchResult returns match result string
func (m *Match) GetMatchResult() string {
	if m.HomeScore == nil || m.AwayScore == nil {
		return "N/A"
	}

	if *m.HomeScore > *m.AwayScore {
		return "Tim Home Menang"
	} else if *m.AwayScore > *m.HomeScore {
		return "Tim Away Menang"
	}
	return "Draw"
}
