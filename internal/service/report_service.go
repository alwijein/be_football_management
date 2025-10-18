package service

import (
	"errors"

	"github.com/alwijein/be_football/internal/dto"
	"github.com/alwijein/be_football/internal/repository"
	"gorm.io/gorm"
)

// ReportService handles report business logic
type ReportService interface {
	GetMatchReport(matchID uint) (*dto.MatchReportResponse, error)
	GetAllMatchReports(pagination *dto.Pagination) ([]dto.MatchReportResponse, *dto.Pagination, error)
	GetDashboardStats() (*dto.DashboardStatsResponse, error)
	GetTeamMatchReports(teamID uint) ([]dto.MatchReportResponse, error)
}

type reportService struct {
	matchRepo  repository.MatchRepository
	goalRepo   repository.GoalRepository
	teamRepo   repository.TeamRepository
	playerRepo repository.PlayerRepository
}

// NewReportService creates new report service
func NewReportService(
	matchRepo repository.MatchRepository,
	goalRepo repository.GoalRepository,
	teamRepo repository.TeamRepository,
	playerRepo repository.PlayerRepository,
) ReportService {
	return &reportService{
		matchRepo:  matchRepo,
		goalRepo:   goalRepo,
		teamRepo:   teamRepo,
		playerRepo: playerRepo,
	}
}

func (s *reportService) GetMatchReport(matchID uint) (*dto.MatchReportResponse, error) {
	// Get match with relations
	match, err := s.matchRepo.FindByIDWithRelations(matchID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("match not found")
		}
		return nil, err
	}

	// Get top scorer
	topScorer, _ := s.goalRepo.GetTopScorerInMatch(matchID)

	// Get team wins count
	homeTeamWins, _ := s.matchRepo.CountTeamWins(match.HomeTeamID, matchID)
	awayTeamWins, _ := s.matchRepo.CountTeamWins(match.AwayTeamID, matchID)

	// Get all goals
	goals, _ := s.goalRepo.FindByMatchID(matchID)

	// Convert to response
	report := &dto.MatchReportResponse{
		ID:           match.ID,
		MatchDate:    match.MatchDate,
		MatchTime:    match.MatchTime,
		HomeScore:    match.HomeScore,
		AwayScore:    match.AwayScore,
		MatchResult:  match.GetMatchResult(),
		HomeTeamWins: homeTeamWins,
		AwayTeamWins: awayTeamWins,
		TopScorer:    topScorer,
		Goals:        make([]dto.GoalResponse, 0),
	}

	// Set teams
	if match.HomeTeam != nil {
		report.HomeTeam = &dto.TeamResponse{
			ID:              match.HomeTeam.ID,
			Name:            match.HomeTeam.Name,
			Logo:            match.HomeTeam.Logo,
			EstablishedYear: match.HomeTeam.EstablishedYear,
			Address:         match.HomeTeam.Address,
			City:            match.HomeTeam.City,
			CreatedAt:       match.HomeTeam.CreatedAt,
			UpdatedAt:       match.HomeTeam.UpdatedAt,
		}
	}

	if match.AwayTeam != nil {
		report.AwayTeam = &dto.TeamResponse{
			ID:              match.AwayTeam.ID,
			Name:            match.AwayTeam.Name,
			Logo:            match.AwayTeam.Logo,
			EstablishedYear: match.AwayTeam.EstablishedYear,
			Address:         match.AwayTeam.Address,
			City:            match.AwayTeam.City,
			CreatedAt:       match.AwayTeam.CreatedAt,
			UpdatedAt:       match.AwayTeam.UpdatedAt,
		}
	}

	// Convert goals
	for _, goal := range goals {
		goalResponse := dto.GoalResponse{
			ID:        goal.ID,
			MatchID:   goal.MatchID,
			PlayerID:  goal.PlayerID,
			Minute:    goal.Minute,
			CreatedAt: goal.CreatedAt,
		}

		if goal.Player != nil {
			goalResponse.Player = &dto.PlayerResponse{
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

			if goal.Player.Team != nil {
				goalResponse.Player.Team = &dto.TeamResponse{
					ID:              goal.Player.Team.ID,
					Name:            goal.Player.Team.Name,
					Logo:            goal.Player.Team.Logo,
					EstablishedYear: goal.Player.Team.EstablishedYear,
					Address:         goal.Player.Team.Address,
					City:            goal.Player.Team.City,
					CreatedAt:       goal.Player.Team.CreatedAt,
					UpdatedAt:       goal.Player.Team.UpdatedAt,
				}
			}
		}

		report.Goals = append(report.Goals, goalResponse)
	}

	return report, nil
}

func (s *reportService) GetAllMatchReports(pagination *dto.Pagination) ([]dto.MatchReportResponse, *dto.Pagination, error) {
	// Get all matches
	matches, err := s.matchRepo.FindAll(pagination)
	if err != nil {
		return nil, nil, err
	}

	total, err := s.matchRepo.Count()
	if err != nil {
		return nil, nil, err
	}

	pagination.Total = total
	pagination.CalculateTotalPages()

	reports := make([]dto.MatchReportResponse, 0)
	for _, match := range matches {
		report, err := s.GetMatchReport(match.ID)
		if err != nil {
			continue
		}
		reports = append(reports, *report)
	}

	return reports, pagination, nil
}

func (s *reportService) GetDashboardStats() (*dto.DashboardStatsResponse, error) {
	// Count total teams
	totalTeams, err := s.teamRepo.Count()
	if err != nil {
		return nil, err
	}

	// Count total matches (schedules)
	totalSchedules, err := s.matchRepo.Count()
	if err != nil {
		return nil, err
	}

	// Count total players
	totalPlayers, err := s.playerRepo.Count()
	if err != nil {
		return nil, err
	}

	return &dto.DashboardStatsResponse{
		TotalTeams:     totalTeams,
		TotalSchedules: totalSchedules,
		TotalPlayers:   totalPlayers,
	}, nil
}

func (s *reportService) GetTeamMatchReports(teamID uint) ([]dto.MatchReportResponse, error) {
	// Check if team exists
	_, err := s.teamRepo.FindByID(teamID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("team not found")
		}
		return nil, err
	}

	// Get all matches for the team
	matches, err := s.matchRepo.FindByTeamID(teamID)
	if err != nil {
		return nil, err
	}

	reports := make([]dto.MatchReportResponse, 0)
	for _, match := range matches {
		report, err := s.GetMatchReport(match.ID)
		if err != nil {
			continue
		}
		reports = append(reports, *report)
	}

	return reports, nil
}
