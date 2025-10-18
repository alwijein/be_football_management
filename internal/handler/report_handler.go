package handler

import (
	"net/http"
	"strconv"

	"github.com/alwijein/be_football/internal/dto"
	"github.com/alwijein/be_football/internal/service"
	"github.com/alwijein/be_football/internal/utils"
	"github.com/gin-gonic/gin"
)

// ReportHandler handles report requests
type ReportHandler struct {
	reportService service.ReportService
}

// NewReportHandler creates new report handler
func NewReportHandler(reportService service.ReportService) *ReportHandler {
	return &ReportHandler{
		reportService: reportService,
	}
}

// convertLogoURLs converts relative logo paths to full URLs in match report
func (h *ReportHandler) convertLogoURLs(c *gin.Context, report *dto.MatchReportResponse) {
	if report.HomeTeam != nil {
		report.HomeTeam.Logo = utils.GetImageURL(c, report.HomeTeam.Logo)
	}
	if report.AwayTeam != nil {
		report.AwayTeam.Logo = utils.GetImageURL(c, report.AwayTeam.Logo)
	}
	for i := range report.Goals {
		if report.Goals[i].Player != nil && report.Goals[i].Player.Team != nil {
			report.Goals[i].Player.Team.Logo = utils.GetImageURL(c, report.Goals[i].Player.Team.Logo)
		}
	}
}

// GetMatchReport handles get match report by ID request
func (h *ReportHandler) GetMatchReport(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid match ID")
		return
	}

	report, err := h.reportService.GetMatchReport(uint(id))
	if err != nil {
		utils.NotFoundResponse(c, err.Error())
		return
	}

	// Convert logo paths to full URLs
	h.convertLogoURLs(c, report)

	utils.SuccessResponse(c, http.StatusOK, "Match report retrieved successfully", report)
}

// GetAllMatchReports handles get all match reports request
func (h *ReportHandler) GetAllMatchReports(c *gin.Context) {
	var pagination dto.Pagination
	if err := c.ShouldBindQuery(&pagination); err != nil {
		pagination.Page = 1
		pagination.Limit = 10
	}

	reports, pag, err := h.reportService.GetAllMatchReports(&pagination)
	if err != nil {
		utils.InternalServerErrorResponse(c, err.Error())
		return
	}

	// Convert logo paths to full URLs for all reports
	for i := range reports {
		h.convertLogoURLs(c, &reports[i])
	}

	utils.SuccessResponseWithPagination(c, "Match reports retrieved successfully", reports, pag)
}

// GetDashboardStats handles get dashboard statistics request
func (h *ReportHandler) GetDashboardStats(c *gin.Context) {
	stats, err := h.reportService.GetDashboardStats()
	if err != nil {
		utils.InternalServerErrorResponse(c, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Dashboard stats retrieved successfully", stats)
}

// GetTeamMatchReports handles get team match reports request
func (h *ReportHandler) GetTeamMatchReports(c *gin.Context) {
	teamID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid team ID")
		return
	}

	reports, err := h.reportService.GetTeamMatchReports(uint(teamID))
	if err != nil {
		utils.NotFoundResponse(c, err.Error())
		return
	}

	// Convert logo paths to full URLs for all reports
	for i := range reports {
		h.convertLogoURLs(c, &reports[i])
	}

	utils.SuccessResponse(c, http.StatusOK, "Team match reports retrieved successfully", reports)
}
