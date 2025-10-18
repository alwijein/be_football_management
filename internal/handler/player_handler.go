package handler

import (
	"net/http"
	"strconv"

	"github.com/alwijein/be_football/internal/dto"
	"github.com/alwijein/be_football/internal/service"
	"github.com/alwijein/be_football/internal/utils"
	"github.com/gin-gonic/gin"
)

// PlayerHandler handles player requests
type PlayerHandler struct {
	playerService service.PlayerService
}

// NewPlayerHandler creates new player handler
func NewPlayerHandler(playerService service.PlayerService) *PlayerHandler {
	return &PlayerHandler{
		playerService: playerService,
	}
}

// Create handles create player request
func (h *PlayerHandler) Create(c *gin.Context) {
	// Get team ID from URL parameter
	teamID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid team ID")
		return
	}

	var req dto.CreatePlayerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err.Error())
		return
	}

	// Set team ID from URL parameter
	teamIDUint := uint(teamID)
	req.TeamID = &teamIDUint

	player, err := h.playerService.Create(&req)
	if err != nil {
		utils.BadRequestResponse(c, err.Error())
		return
	}

	response := dto.PlayerResponse{
		ID:           player.ID,
		TeamID:       player.TeamID,
		Name:         player.Name,
		Height:       player.Height,
		Weight:       player.Weight,
		Position:     string(player.Position),
		JerseyNumber: player.JerseyNumber,
		CreatedAt:    player.CreatedAt,
		UpdatedAt:    player.UpdatedAt,
	}

	if player.Team != nil {
		response.Team = &dto.TeamResponse{
			ID:              player.Team.ID,
			Name:            player.Team.Name,
			Logo:            utils.GetImageURL(c, player.Team.Logo),
			EstablishedYear: player.Team.EstablishedYear,
			Address:         player.Team.Address,
			City:            player.Team.City,
			CreatedAt:       player.Team.CreatedAt,
			UpdatedAt:       player.Team.UpdatedAt,
		}
	}

	utils.SuccessResponse(c, http.StatusCreated, "Player created successfully", response)
}

// GetByID handles get player by ID request
func (h *PlayerHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid player ID")
		return
	}

	player, err := h.playerService.GetByID(uint(id))
	if err != nil {
		utils.NotFoundResponse(c, err.Error())
		return
	}

	response := dto.PlayerResponse{
		ID:           player.ID,
		TeamID:       player.TeamID,
		Name:         player.Name,
		Height:       player.Height,
		Weight:       player.Weight,
		Position:     string(player.Position),
		JerseyNumber: player.JerseyNumber,
		CreatedAt:    player.CreatedAt,
		UpdatedAt:    player.UpdatedAt,
	}

	if player.Team != nil {
		response.Team = &dto.TeamResponse{
			ID:              player.Team.ID,
			Name:            player.Team.Name,
			Logo:            utils.GetImageURL(c, player.Team.Logo),
			EstablishedYear: player.Team.EstablishedYear,
			Address:         player.Team.Address,
			City:            player.Team.City,
			CreatedAt:       player.Team.CreatedAt,
			UpdatedAt:       player.Team.UpdatedAt,
		}
	}

	utils.SuccessResponse(c, http.StatusOK, "Player retrieved successfully", response)
}

// GetAll handles get all players request
func (h *PlayerHandler) GetAll(c *gin.Context) {
	var pagination dto.Pagination
	if err := c.ShouldBindQuery(&pagination); err != nil {
		pagination.Page = 1
		pagination.Limit = 10
	}

	players, pag, err := h.playerService.GetAll(&pagination)
	if err != nil {
		utils.InternalServerErrorResponse(c, err.Error())
		return
	}

	var response []dto.PlayerResponse
	for _, player := range players {
		playerResp := dto.PlayerResponse{
			ID:           player.ID,
			TeamID:       player.TeamID,
			Name:         player.Name,
			Height:       player.Height,
			Weight:       player.Weight,
			Position:     string(player.Position),
			JerseyNumber: player.JerseyNumber,
			CreatedAt:    player.CreatedAt,
			UpdatedAt:    player.UpdatedAt,
		}

		if player.Team != nil {
			playerResp.Team = &dto.TeamResponse{
				ID:              player.Team.ID,
				Name:            player.Team.Name,
				Logo:            utils.GetImageURL(c, player.Team.Logo),
				EstablishedYear: player.Team.EstablishedYear,
				Address:         player.Team.Address,
				City:            player.Team.City,
				CreatedAt:       player.Team.CreatedAt,
				UpdatedAt:       player.Team.UpdatedAt,
			}
		}

		response = append(response, playerResp)
	}

	utils.SuccessResponseWithPagination(c, "Players retrieved successfully", response, pag)
}

// GetByTeamID handles get players by team ID request
func (h *PlayerHandler) GetByTeamID(c *gin.Context) {
	teamID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid team ID")
		return
	}

	players, err := h.playerService.GetByTeamID(uint(teamID))
	if err != nil {
		utils.BadRequestResponse(c, err.Error())
		return
	}

	var response []dto.PlayerResponse
	for _, player := range players {
		response = append(response, dto.PlayerResponse{
			ID:           player.ID,
			TeamID:       player.TeamID,
			Name:         player.Name,
			Height:       player.Height,
			Weight:       player.Weight,
			Position:     string(player.Position),
			JerseyNumber: player.JerseyNumber,
			CreatedAt:    player.CreatedAt,
			UpdatedAt:    player.UpdatedAt,
		})
	}

	utils.SuccessResponse(c, http.StatusOK, "Players retrieved successfully", response)
}

// Update handles update player request
func (h *PlayerHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid player ID")
		return
	}

	var req dto.UpdatePlayerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err.Error())
		return
	}

	player, err := h.playerService.Update(uint(id), &req)
	if err != nil {
		utils.BadRequestResponse(c, err.Error())
		return
	}

	response := dto.PlayerResponse{
		ID:           player.ID,
		TeamID:       player.TeamID,
		Name:         player.Name,
		Height:       player.Height,
		Weight:       player.Weight,
		Position:     string(player.Position),
		JerseyNumber: player.JerseyNumber,
		CreatedAt:    player.CreatedAt,
		UpdatedAt:    player.UpdatedAt,
	}

	if player.Team != nil {
		response.Team = &dto.TeamResponse{
			ID:              player.Team.ID,
			Name:            player.Team.Name,
			Logo:            utils.GetImageURL(c, player.Team.Logo),
			EstablishedYear: player.Team.EstablishedYear,
			Address:         player.Team.Address,
			City:            player.Team.City,
			CreatedAt:       player.Team.CreatedAt,
			UpdatedAt:       player.Team.UpdatedAt,
		}
	}

	utils.SuccessResponse(c, http.StatusOK, "Player updated successfully", response)
}

// Delete handles delete player request
func (h *PlayerHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid player ID")
		return
	}

	if err := h.playerService.Delete(uint(id)); err != nil {
		utils.BadRequestResponse(c, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Player deleted successfully", nil)
}
