package service

import (
	"errors"
	"time"

	"github.com/alwijein/be_football/internal/domain"
	"github.com/alwijein/be_football/internal/dto"
	"github.com/alwijein/be_football/internal/repository"
	"gorm.io/gorm"
)

// MatchService handles match business logic
type MatchService interface {
	Create(req *dto.CreateMatchRequest) (*domain.Match, error)
	GetByID(id uint) (*domain.Match, error)
	GetByIDWithRelations(id uint) (*domain.Match, error)
	GetAll(pagination *dto.Pagination) ([]domain.Match, *dto.Pagination, error)
	GetTodayMatches() ([]domain.Match, error)
	GetMatchScorers(matchID uint) ([]dto.GoalResponse, error)
	Update(id uint, req *dto.UpdateMatchRequest) (*domain.Match, error)
	Delete(id uint) error
	SetResult(matchID uint, req *dto.MatchResultRequest) (*domain.Match, error)
	AddGoal(matchID uint, req *dto.CreateGoalRequest) (*domain.Goal, error)
	DeleteGoal(goalID uint) error
}

type matchService struct {
	matchRepo  repository.MatchRepository
	teamRepo   repository.TeamRepository
	playerRepo repository.PlayerRepository
	goalRepo   repository.GoalRepository
}

// NewMatchService creates new match service
func NewMatchService(
	matchRepo repository.MatchRepository,
	teamRepo repository.TeamRepository,
	playerRepo repository.PlayerRepository,
	goalRepo repository.GoalRepository,
) MatchService {
	return &matchService{
		matchRepo:  matchRepo,
		teamRepo:   teamRepo,
		playerRepo: playerRepo,
		goalRepo:   goalRepo,
	}
}

func (s *matchService) Create(req *dto.CreateMatchRequest) (*domain.Match, error) {
	// Validate teams are different
	if req.HomeTeamID == req.AwayTeamID {
		return nil, errors.New("home team and away team must be different")
	}

	// Check if both teams exist
	_, err := s.teamRepo.FindByID(req.HomeTeamID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("home team not found")
		}
		return nil, err
	}

	_, err = s.teamRepo.FindByID(req.AwayTeamID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("away team not found")
		}
		return nil, err
	}

	// Parse match date
	matchDate, err := time.Parse("2006-01-02", req.MatchDate)
	if err != nil {
		return nil, errors.New("invalid match date format, use YYYY-MM-DD")
	}

	match := &domain.Match{
		HomeTeamID: req.HomeTeamID,
		AwayTeamID: req.AwayTeamID,
		MatchDate:  matchDate,
		MatchTime:  req.MatchTime,
		Status:     domain.StatusScheduled,
	}

	if err := s.matchRepo.Create(match); err != nil {
		return nil, err
	}

	// Load relations
	match, _ = s.matchRepo.FindByIDWithRelations(match.ID)
	return match, nil
}

func (s *matchService) GetByID(id uint) (*domain.Match, error) {
	match, err := s.matchRepo.FindByIDWithRelations(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("match not found")
		}
		return nil, err
	}
	return match, nil
}

func (s *matchService) GetByIDWithRelations(id uint) (*domain.Match, error) {
	return s.GetByID(id)
}

func (s *matchService) GetAll(pagination *dto.Pagination) ([]domain.Match, *dto.Pagination, error) {
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

	return matches, pagination, nil
}

func (s *matchService) Update(id uint, req *dto.UpdateMatchRequest) (*domain.Match, error) {
	match, err := s.matchRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("match not found")
		}
		return nil, err
	}

	// Don't allow updating completed matches
	if match.Status == domain.StatusCompleted {
		return nil, errors.New("cannot update completed match")
	}

	// Validate and update teams
	if req.HomeTeamID > 0 && req.AwayTeamID > 0 {
		if req.HomeTeamID == req.AwayTeamID {
			return nil, errors.New("home team and away team must be different")
		}
	}

	if req.HomeTeamID > 0 {
		_, err := s.teamRepo.FindByID(req.HomeTeamID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("home team not found")
			}
			return nil, err
		}
		match.HomeTeamID = req.HomeTeamID
	}

	if req.AwayTeamID > 0 {
		_, err := s.teamRepo.FindByID(req.AwayTeamID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("away team not found")
			}
			return nil, err
		}
		match.AwayTeamID = req.AwayTeamID
	}

	if req.MatchDate != "" {
		matchDate, err := time.Parse("2006-01-02", req.MatchDate)
		if err != nil {
			return nil, errors.New("invalid match date format, use YYYY-MM-DD")
		}
		match.MatchDate = matchDate
	}

	if req.MatchTime != "" {
		match.MatchTime = req.MatchTime
	}

	if err := s.matchRepo.Update(match); err != nil {
		return nil, err
	}

	match, _ = s.matchRepo.FindByIDWithRelations(match.ID)
	return match, nil
}

func (s *matchService) Delete(id uint) error {
	match, err := s.matchRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("match not found")
		}
		return err
	}

	return s.matchRepo.Delete(match.ID)
}

func (s *matchService) SetResult(matchID uint, req *dto.MatchResultRequest) (*domain.Match, error) {
	match, err := s.matchRepo.FindByID(matchID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("match not found")
		}
		return nil, err
	}

	match.HomeScore = req.HomeScore
	match.AwayScore = req.AwayScore

	// Set status if provided, otherwise default to Completed
	if req.Status != "" {
		match.Status = domain.MatchStatus(req.Status)
	} else {
		match.Status = domain.StatusCompleted
	}

	if err := s.matchRepo.Update(match); err != nil {
		return nil, err
	}

	match, _ = s.matchRepo.FindByIDWithRelations(match.ID)
	return match, nil
}

func (s *matchService) AddGoal(matchID uint, req *dto.CreateGoalRequest) (*domain.Goal, error) {
	// Check if match exists
	match, err := s.matchRepo.FindByID(matchID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("match not found")
		}
		return nil, err
	}

	// Check if player exists
	player, err := s.playerRepo.FindByIDWithTeam(req.PlayerID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("player not found")
		}
		return nil, err
	}

	// Validate player is from one of the teams
	if player.TeamID == nil {
		return nil, errors.New("player must be assigned to a team")
	}
	if *player.TeamID != match.HomeTeamID && *player.TeamID != match.AwayTeamID {
		return nil, errors.New("player is not from either team in this match")
	}

	goal := &domain.Goal{
		MatchID:  matchID,
		PlayerID: req.PlayerID,
		Minute:   req.Minute,
	}

	if err := s.goalRepo.Create(goal); err != nil {
		return nil, err
	}

	goal, _ = s.goalRepo.FindByID(goal.ID)
	return goal, nil
}

func (s *matchService) DeleteGoal(goalID uint) error {
	goal, err := s.goalRepo.FindByID(goalID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("goal not found")
		}
		return err
	}

	return s.goalRepo.Delete(goal.ID)
}

func (s *matchService) GetTodayMatches() ([]domain.Match, error) {
	// Get today's date
	today := time.Now().Format("2006-01-02")
	todayDate, _ := time.Parse("2006-01-02", today)

	// Get all matches
	var pagination dto.Pagination
	pagination.Page = 1
	pagination.Limit = 100 // Get all matches

	matches, err := s.matchRepo.FindAll(&pagination)
	if err != nil {
		return nil, err
	}

	// Filter today's matches
	todayMatches := make([]domain.Match, 0)
	for _, match := range matches {
		if match.MatchDate.Format("2006-01-02") == todayDate.Format("2006-01-02") {
			// Load relations
			loadedMatch, err := s.matchRepo.FindByIDWithRelations(match.ID)
			if err == nil {
				todayMatches = append(todayMatches, *loadedMatch)
			}
		}
	}

	return todayMatches, nil
}

func (s *matchService) GetMatchScorers(matchID uint) ([]dto.GoalResponse, error) {
	// Check if match exists
	match, err := s.matchRepo.FindByIDWithRelations(matchID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("match not found")
		}
		return nil, err
	}

	// Get all goals for this match
	goals := match.Goals

	// Convert to GoalResponse
	goalResponses := make([]dto.GoalResponse, 0)
	for _, goal := range goals {
		goalResponses = append(goalResponses, dto.GoalResponse{
			ID:      goal.ID,
			MatchID: goal.MatchID,
			Player: &dto.PlayerResponse{
				ID:           goal.Player.ID,
				TeamID:       goal.Player.TeamID,
				Name:         goal.Player.Name,
				Position:     string(goal.Player.Position),
				JerseyNumber: goal.Player.JerseyNumber,
				Team: &dto.TeamResponse{
					ID:   goal.Player.Team.ID,
					Name: goal.Player.Team.Name,
					Logo: goal.Player.Team.Logo,
				},
			},
			Minute:    goal.Minute,
			CreatedAt: goal.CreatedAt,
		})
	}

	return goalResponses, nil
}
