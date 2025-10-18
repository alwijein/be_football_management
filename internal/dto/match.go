package dto

import "time"

// CreateMatchRequest represents create match request
type CreateMatchRequest struct {
	HomeTeamID uint   `json:"home_team_id" binding:"required"`
	AwayTeamID uint   `json:"away_team_id" binding:"required"`
	MatchDate  string `json:"match_date" binding:"required"` // Format: YYYY-MM-DD
	MatchTime  string `json:"match_time" binding:"required,len=5"` // Format: HH:MM
}

// UpdateMatchRequest represents update match request
type UpdateMatchRequest struct {
	HomeTeamID uint   `json:"home_team_id" binding:"omitempty"`
	AwayTeamID uint   `json:"away_team_id" binding:"omitempty"`
	MatchDate  string `json:"match_date" binding:"omitempty"`
	MatchTime  string `json:"match_time" binding:"omitempty,len=5"`
}

// MatchResultRequest represents match result request
type MatchResultRequest struct {
	HomeScore int `json:"home_score" binding:"required,min=0"`
	AwayScore int `json:"away_score" binding:"required,min=0"`
}

// MatchResponse represents match response
type MatchResponse struct {
	ID         uint          `json:"id"`
	HomeTeamID uint          `json:"home_team_id"`
	AwayTeamID uint          `json:"away_team_id"`
	MatchDate  time.Time     `json:"match_date"`
	MatchTime  string        `json:"match_time"`
	HomeScore  *int          `json:"home_score"`
	AwayScore  *int          `json:"away_score"`
	Status     string        `json:"status"`
	HomeTeam   *TeamResponse `json:"home_team,omitempty"`
	AwayTeam   *TeamResponse `json:"away_team,omitempty"`
	CreatedAt  time.Time     `json:"created_at"`
	UpdatedAt  time.Time     `json:"updated_at"`
}

// CreateGoalRequest represents create goal request
type CreateGoalRequest struct {
	PlayerID uint `json:"player_id" binding:"required"`
	Minute   int  `json:"minute" binding:"required,min=1,max=120"`
}

// GoalResponse represents goal response
type GoalResponse struct {
	ID        uint            `json:"id"`
	MatchID   uint            `json:"match_id"`
	PlayerID  uint            `json:"player_id"`
	Minute    int             `json:"minute"`
	Player    *PlayerResponse `json:"player,omitempty"`
	CreatedAt time.Time       `json:"created_at"`
}
