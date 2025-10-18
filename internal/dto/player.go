package dto

import "time"

// CreatePlayerRequest represents create player request
type CreatePlayerRequest struct {
	TeamID       *uint  `json:"team_id" binding:"omitempty"`
	Name         string `json:"name" binding:"required,min=3,max=100"`
	Height       int    `json:"height" binding:"required,min=100,max=250"`
	Weight       int    `json:"weight" binding:"required,min=30,max=200"`
	Position     string `json:"position" binding:"required"`
	JerseyNumber int    `json:"jersey_number" binding:"required,min=1,max=99"`
}

// UpdatePlayerRequest represents update player request
type UpdatePlayerRequest struct {
	TeamID       *uint  `json:"team_id" binding:"omitempty"`
	Name         string `json:"name" binding:"omitempty,min=3,max=100"`
	Height       int    `json:"height" binding:"omitempty,min=100,max=250"`
	Weight       int    `json:"weight" binding:"omitempty,min=30,max=200"`
	Position     string `json:"position" binding:"omitempty"`
	JerseyNumber int    `json:"jersey_number" binding:"omitempty,min=1,max=99"`
}

// PlayerResponse represents player response
type PlayerResponse struct {
	ID           uint          `json:"id"`
	TeamID       *uint         `json:"team_id"`
	Name         string        `json:"name"`
	Height       int           `json:"height"`
	Weight       int           `json:"weight"`
	Position     string        `json:"position"`
	JerseyNumber int           `json:"jersey_number"`
	Team         *TeamResponse `json:"team,omitempty"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
}
