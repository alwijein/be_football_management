package handler

import (
	"net/http"
	"strconv"

	"github.com/alwijein/be_football/internal/dto"
	"github.com/alwijein/be_football/internal/service"
	"github.com/alwijein/be_football/internal/utils"
	"github.com/gin-gonic/gin"
)

// TeamHandler handles team requests
type TeamHandler struct {
	teamService service.TeamService
}

// NewTeamHandler creates new team handler
func NewTeamHandler(teamService service.TeamService) *TeamHandler {
	return &TeamHandler{
		teamService: teamService,
	}
}

// Create handles create team request
func (h *TeamHandler) Create(c *gin.Context) {
	// Parse form data
	var req dto.CreateTeamRequest

	// Get form values
	req.Name = c.PostForm("name")
	req.Address = c.PostForm("address")
	req.City = c.PostForm("city")

	// Parse established_year
	year, err := strconv.Atoi(c.PostForm("established_year"))
	if err != nil {
		utils.ValidationErrorResponse(c, "Invalid established_year: must be a number")
		return
	}
	req.EstablishedYear = year

	// Validate required fields
	if req.Name == "" || req.Address == "" || req.City == "" {
		utils.ValidationErrorResponse(c, "name, address, and city are required")
		return
	}

	// Validate year range
	if year < 1800 || year > 2100 {
		utils.ValidationErrorResponse(c, "established_year must be between 1800 and 2100")
		return
	}

	// Handle logo upload
	logoPath := ""
	_, err = c.FormFile("logo")
	if err == nil {
		// File was provided, upload it
		logoPath, err = utils.UploadImage(c, "logo", "teams")
		if err != nil {
			utils.ValidationErrorResponse(c, "Logo upload failed: "+err.Error())
			return
		}
	}
	req.Logo = logoPath

	team, err := h.teamService.Create(&req)
	if err != nil {
		// Delete uploaded file if team creation fails
		if logoPath != "" {
			utils.DeleteImage(logoPath)
		}
		utils.BadRequestResponse(c, err.Error())
		return
	}

	response := dto.TeamResponse{
		ID:              team.ID,
		Name:            team.Name,
		Logo:            utils.GetImageURL(c, team.Logo),
		EstablishedYear: team.EstablishedYear,
		Address:         team.Address,
		City:            team.City,
		CreatedAt:       team.CreatedAt,
		UpdatedAt:       team.UpdatedAt,
	}

	utils.SuccessResponse(c, http.StatusCreated, "Team created successfully", response)
}

// GetByID handles get team by ID request
func (h *TeamHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid team ID")
		return
	}

	team, err := h.teamService.GetByIDWithPlayers(uint(id))
	if err != nil {
		utils.NotFoundResponse(c, err.Error())
		return
	}

	response := dto.TeamWithPlayersResponse{
		TeamResponse: dto.TeamResponse{
			ID:              team.ID,
			Name:            team.Name,
			Logo:            utils.GetImageURL(c, team.Logo),
			EstablishedYear: team.EstablishedYear,
			Address:         team.Address,
			City:            team.City,
			CreatedAt:       team.CreatedAt,
			UpdatedAt:       team.UpdatedAt,
		},
		Players: make([]dto.PlayerResponse, 0),
	}

	for _, player := range team.Players {
		response.Players = append(response.Players, dto.PlayerResponse{
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

	utils.SuccessResponse(c, http.StatusOK, "Team retrieved successfully", response)
}

// GetRegistered handles get registered teams request (for homepage)
func (h *TeamHandler) GetRegistered(c *gin.Context) {
	// Default limit 5, can be overridden via query param
	limit := 5
	if limitStr := c.Query("limit"); limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	teams, err := h.teamService.GetRegistered(limit)
	if err != nil {
		utils.InternalServerErrorResponse(c, err.Error())
		return
	}

	var response []dto.TeamResponse
	for _, team := range teams {
		response = append(response, dto.TeamResponse{
			ID:              team.ID,
			Name:            team.Name,
			Logo:            utils.GetImageURL(c, team.Logo),
			EstablishedYear: team.EstablishedYear,
			Address:         team.Address,
			City:            team.City,
			CreatedAt:       team.CreatedAt,
			UpdatedAt:       team.UpdatedAt,
		})
	}

	utils.SuccessResponse(c, http.StatusOK, "Registered teams retrieved successfully", response)
}

// GetAll handles get all teams request
func (h *TeamHandler) GetAll(c *gin.Context) {
	var pagination dto.Pagination
	if err := c.ShouldBindQuery(&pagination); err != nil {
		pagination.Page = 1
		pagination.Limit = 10
	}

	teams, pag, err := h.teamService.GetAll(&pagination)
	if err != nil {
		utils.InternalServerErrorResponse(c, err.Error())
		return
	}

	var response []dto.TeamResponse
	for _, team := range teams {
		response = append(response, dto.TeamResponse{
			ID:              team.ID,
			Name:            team.Name,
			Logo:            utils.GetImageURL(c, team.Logo),
			EstablishedYear: team.EstablishedYear,
			Address:         team.Address,
			City:            team.City,
			CreatedAt:       team.CreatedAt,
			UpdatedAt:       team.UpdatedAt,
		})
	}

	utils.SuccessResponseWithPagination(c, "Teams retrieved successfully", response, pag)
}

// Update handles update team request
func (h *TeamHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid team ID")
		return
	}

	// Get existing team first
	existingTeam, err := h.teamService.GetByID(uint(id))
	if err != nil {
		utils.NotFoundResponse(c, "Team not found")
		return
	}

	// Parse form data
	var req dto.UpdateTeamRequest

	// Get form values (all optional for update)
	if name := c.PostForm("name"); name != "" {
		req.Name = name
	}
	if address := c.PostForm("address"); address != "" {
		req.Address = address
	}
	if city := c.PostForm("city"); city != "" {
		req.City = city
	}
	if yearStr := c.PostForm("established_year"); yearStr != "" {
		year, err := strconv.Atoi(yearStr)
		if err != nil {
			utils.ValidationErrorResponse(c, "Invalid established_year: must be a number")
			return
		}
		if year < 1800 || year > 2100 {
			utils.ValidationErrorResponse(c, "established_year must be between 1800 and 2100")
			return
		}
		req.EstablishedYear = year
	}

	// Handle logo upload if provided
	_, err = c.FormFile("logo")
	if err == nil {
		// New logo file provided, upload it
		logoPath, err := utils.UploadImage(c, "logo", "teams")
		if err != nil {
			utils.ValidationErrorResponse(c, "Logo upload failed: "+err.Error())
			return
		}

		// Delete old logo if exists
		if existingTeam.Logo != "" {
			utils.DeleteImage(existingTeam.Logo)
		}

		req.Logo = logoPath
	}

	team, err := h.teamService.Update(uint(id), &req)
	if err != nil {
		utils.BadRequestResponse(c, err.Error())
		return
	}

	response := dto.TeamResponse{
		ID:              team.ID,
		Name:            team.Name,
		Logo:            utils.GetImageURL(c, team.Logo),
		EstablishedYear: team.EstablishedYear,
		Address:         team.Address,
		City:            team.City,
		CreatedAt:       team.CreatedAt,
		UpdatedAt:       team.UpdatedAt,
	}

	utils.SuccessResponse(c, http.StatusOK, "Team updated successfully", response)
}

// Delete handles delete team request
func (h *TeamHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid team ID")
		return
	}

	// Get team first to delete logo file
	team, err := h.teamService.GetByID(uint(id))
	if err != nil {
		utils.NotFoundResponse(c, "Team not found")
		return
	}

	// Delete team
	if err := h.teamService.Delete(uint(id)); err != nil {
		utils.BadRequestResponse(c, err.Error())
		return
	}

	// Delete logo file if exists
	if team.Logo != "" {
		utils.DeleteImage(team.Logo)
	}

	utils.SuccessResponse(c, http.StatusOK, "Team deleted successfully", nil)
}
