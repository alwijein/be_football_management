package handler

import (
	"net/http"
	"strconv"

	"github.com/alwijein/be_football/internal/domain"
	"github.com/alwijein/be_football/internal/dto"
	"github.com/alwijein/be_football/internal/service"
	"github.com/alwijein/be_football/internal/utils"
	"github.com/gin-gonic/gin"
)

// MatchHandler handles match requests
type MatchHandler struct {
	matchService service.MatchService
}

// NewMatchHandler creates new match handler
func NewMatchHandler(matchService service.MatchService) *MatchHandler {
	return &MatchHandler{
		matchService: matchService,
	}
}

// Create handles create match request
func (h *MatchHandler) Create(c *gin.Context) {
	var req dto.CreateMatchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err.Error())
		return
	}

	match, err := h.matchService.Create(&req)
	if err != nil {
		utils.BadRequestResponse(c, err.Error())
		return
	}

	response := h.toMatchResponse(match, c)
	utils.SuccessResponse(c, http.StatusCreated, "Match created successfully", response)
}

// GetByID handles get match by ID request
func (h *MatchHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid match ID")
		return
	}

	match, err := h.matchService.GetByID(uint(id))
	if err != nil {
		utils.NotFoundResponse(c, err.Error())
		return
	}

	response := h.toMatchResponse(match, c)
	utils.SuccessResponse(c, http.StatusOK, "Match retrieved successfully", response)
}

// GetAll handles get all matches request
func (h *MatchHandler) GetAll(c *gin.Context) {
	var pagination dto.Pagination
	if err := c.ShouldBindQuery(&pagination); err != nil {
		pagination.Page = 1
		pagination.Limit = 10
	}

	matches, pag, err := h.matchService.GetAll(&pagination)
	if err != nil {
		utils.InternalServerErrorResponse(c, err.Error())
		return
	}

	var response []dto.MatchResponse
	for _, match := range matches {
		response = append(response, h.toMatchResponse(&match, c))
	}

	utils.SuccessResponseWithPagination(c, "Matches retrieved successfully", response, pag)
}

// GetTodayMatches handles get today's matches request
func (h *MatchHandler) GetTodayMatches(c *gin.Context) {
	matches, err := h.matchService.GetTodayMatches()
	if err != nil {
		utils.InternalServerErrorResponse(c, err.Error())
		return
	}

	var response []dto.MatchResponse
	for _, match := range matches {
		response = append(response, h.toMatchResponse(&match, c))
	}

	utils.SuccessResponse(c, http.StatusOK, "Today's matches retrieved successfully", response)
}

// GetScorers handles get match scorers request
func (h *MatchHandler) GetScorers(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid schedule ID")
		return
	}

	scorers, err := h.matchService.GetMatchScorers(uint(id))
	if err != nil {
		utils.NotFoundResponse(c, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Match scorers retrieved successfully", scorers)
}

// Update handles update match request
func (h *MatchHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid match ID")
		return
	}

	var req dto.UpdateMatchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err.Error())
		return
	}

	match, err := h.matchService.Update(uint(id), &req)
	if err != nil {
		utils.BadRequestResponse(c, err.Error())
		return
	}

	response := h.toMatchResponse(match, c)
	utils.SuccessResponse(c, http.StatusOK, "Match updated successfully", response)
}

// Delete handles delete match request
func (h *MatchHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid match ID")
		return
	}

	if err := h.matchService.Delete(uint(id)); err != nil {
		utils.BadRequestResponse(c, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Match deleted successfully", nil)
}

// SetResult handles set match result request
func (h *MatchHandler) SetResult(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid match ID")
		return
	}

	var req dto.MatchResultRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err.Error())
		return
	}

	match, err := h.matchService.SetResult(uint(id), &req)
	if err != nil {
		utils.BadRequestResponse(c, err.Error())
		return
	}

	response := h.toMatchResponse(match, c)
	utils.SuccessResponse(c, http.StatusOK, "Match result set successfully", response)
}

// AddGoal handles add goal to match request
func (h *MatchHandler) AddGoal(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid match ID")
		return
	}

	var req dto.CreateGoalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err.Error())
		return
	}

	goal, err := h.matchService.AddGoal(uint(id), &req)
	if err != nil {
		utils.BadRequestResponse(c, err.Error())
		return
	}

	response := dto.GoalResponse{
		ID:        goal.ID,
		MatchID:   goal.MatchID,
		PlayerID:  goal.PlayerID,
		Minute:    goal.Minute,
		CreatedAt: goal.CreatedAt,
	}

	if goal.Player != nil {
		response.Player = &dto.PlayerResponse{
			ID:           goal.Player.ID,
			TeamID:       goal.Player.TeamID,
			Name:         goal.Player.Name,
			Height:       goal.Player.Height,
			Weight:       goal.Player.Weight,
			Position:     string(goal.Player.Position),
			JerseyNumber: goal.Player.JerseyNumber,
			CreatedAt:    goal.Player.CreatedAt,
			UpdatedAt:    goal.Player.UpdatedAt,
		}
	}

	utils.SuccessResponse(c, http.StatusCreated, "Goal added successfully", response)
}

// DeleteGoal handles delete goal request
func (h *MatchHandler) DeleteGoal(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid goal ID")
		return
	}

	if err := h.matchService.DeleteGoal(uint(id)); err != nil {
		utils.BadRequestResponse(c, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Goal deleted successfully", nil)
}

// Helper function to convert domain Match to MatchResponse
func (h *MatchHandler) toMatchResponse(match *domain.Match, c *gin.Context) dto.MatchResponse {
	response := dto.MatchResponse{
		ID:         match.ID,
		HomeTeamID: match.HomeTeamID,
		AwayTeamID: match.AwayTeamID,
		MatchDate:  match.MatchDate,
		MatchTime:  match.MatchTime,
		HomeScore:  match.HomeScore,
		AwayScore:  match.AwayScore,
		Status:     string(match.Status),
		CreatedAt:  match.CreatedAt,
		UpdatedAt:  match.UpdatedAt,
	}

	if match.HomeTeam != nil {
		response.HomeTeam = &dto.TeamResponse{
			ID:              match.HomeTeam.ID,
			Name:            match.HomeTeam.Name,
			Logo:            utils.GetImageURL(c, match.HomeTeam.Logo),
			EstablishedYear: match.HomeTeam.EstablishedYear,
			Address:         match.HomeTeam.Address,
			City:            match.HomeTeam.City,
			CreatedAt:       match.HomeTeam.CreatedAt,
			UpdatedAt:       match.HomeTeam.UpdatedAt,
		}
	}

	if match.AwayTeam != nil {
		response.AwayTeam = &dto.TeamResponse{
			ID:              match.AwayTeam.ID,
			Name:            match.AwayTeam.Name,
			Logo:            utils.GetImageURL(c, match.AwayTeam.Logo),
			EstablishedYear: match.AwayTeam.EstablishedYear,
			Address:         match.AwayTeam.Address,
			City:            match.AwayTeam.City,
			CreatedAt:       match.AwayTeam.CreatedAt,
			UpdatedAt:       match.AwayTeam.UpdatedAt,
		}
	}

	return response
}
