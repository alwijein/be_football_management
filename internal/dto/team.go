package dto

import "time"

// CreateTeamRequest represents create team request
type CreateTeamRequest struct {
	Name            string `json:"name" binding:"required,min=3,max=100"`
	Logo            string `json:"logo" binding:"omitempty"`
	EstablishedYear int    `json:"established_year" binding:"required,min=1800,max=2100"`
	Address         string `json:"address" binding:"required"`
	City            string `json:"city" binding:"required,min=2,max=100"`
}

// UpdateTeamRequest represents update team request
type UpdateTeamRequest struct {
	Name            string `json:"name" binding:"omitempty,min=3,max=100"`
	Logo            string `json:"logo" binding:"omitempty"`
	EstablishedYear int    `json:"established_year" binding:"omitempty,min=1800,max=2100"`
	Address         string `json:"address" binding:"omitempty"`
	City            string `json:"city" binding:"omitempty,min=2,max=100"`
}

// TeamResponse represents team response
type TeamResponse struct {
	ID              uint      `json:"id"`
	Name            string    `json:"name"`
	Logo            string    `json:"logo"`
	EstablishedYear int       `json:"established_year"`
	Address         string    `json:"address"`
	City            string    `json:"city"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// TeamWithPlayersResponse represents team with players response
type TeamWithPlayersResponse struct {
	TeamResponse
	Players []PlayerResponse `json:"players"`
}
