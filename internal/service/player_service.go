package service

import (
	"errors"

	"github.com/alwijein/be_football/internal/domain"
	"github.com/alwijein/be_football/internal/dto"
	"github.com/alwijein/be_football/internal/repository"
	"gorm.io/gorm"
)

// PlayerService handles player business logic
type PlayerService interface {
	Create(req *dto.CreatePlayerRequest) (*domain.Player, error)
	GetByID(id uint) (*domain.Player, error)
	GetAll(pagination *dto.Pagination) ([]domain.Player, *dto.Pagination, error)
	GetByTeamID(teamID uint) ([]domain.Player, error)
	Update(id uint, req *dto.UpdatePlayerRequest) (*domain.Player, error)
	Delete(id uint) error
}

type playerService struct {
	playerRepo repository.PlayerRepository
	teamRepo   repository.TeamRepository
}

// NewPlayerService creates new player service
func NewPlayerService(playerRepo repository.PlayerRepository, teamRepo repository.TeamRepository) PlayerService {
	return &playerService{
		playerRepo: playerRepo,
		teamRepo:   teamRepo,
	}
}

func (s *playerService) Create(req *dto.CreatePlayerRequest) (*domain.Player, error) {
	// Validate position
	if !domain.IsValidPosition(req.Position) {
		return nil, errors.New("invalid player position")
	}

	// Check if team exists
	if req.TeamID != nil {
		_, err := s.teamRepo.FindByID(*req.TeamID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("team not found")
			}
			return nil, err
		}

		// Check jersey number uniqueness in team
		exists, err := s.playerRepo.ExistsByJerseyNumberInTeam(*req.TeamID, req.JerseyNumber, 0)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, errors.New("jersey number already exists in this team")
		}
	}

	player := &domain.Player{
		TeamID:       req.TeamID,
		Name:         req.Name,
		Height:       req.Height,
		Weight:       req.Weight,
		Position:     domain.PlayerPosition(req.Position),
		JerseyNumber: req.JerseyNumber,
	}

	if err := s.playerRepo.Create(player); err != nil {
		return nil, err
	}

	// Load team relation
	if player.TeamID != nil {
		player, _ = s.playerRepo.FindByIDWithTeam(player.ID)
	}

	return player, nil
}

func (s *playerService) GetByID(id uint) (*domain.Player, error) {
	player, err := s.playerRepo.FindByIDWithTeam(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("player not found")
		}
		return nil, err
	}
	return player, nil
}

func (s *playerService) GetAll(pagination *dto.Pagination) ([]domain.Player, *dto.Pagination, error) {
	players, err := s.playerRepo.FindAll(pagination)
	if err != nil {
		return nil, nil, err
	}

	total, err := s.playerRepo.Count()
	if err != nil {
		return nil, nil, err
	}

	pagination.Total = total
	pagination.CalculateTotalPages()

	return players, pagination, nil
}

func (s *playerService) GetByTeamID(teamID uint) ([]domain.Player, error) {
	// Check if team exists
	_, err := s.teamRepo.FindByID(teamID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("team not found")
		}
		return nil, err
	}

	return s.playerRepo.FindByTeamID(teamID)
}

func (s *playerService) Update(id uint, req *dto.UpdatePlayerRequest) (*domain.Player, error) {
	player, err := s.playerRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("player not found")
		}
		return nil, err
	}

	// Validate position if provided
	if req.Position != "" && !domain.IsValidPosition(req.Position) {
		return nil, errors.New("invalid player position")
	}

	// Check if team exists and jersey number is unique
	if req.TeamID != nil {
		_, err := s.teamRepo.FindByID(*req.TeamID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("team not found")
			}
			return nil, err
		}

		// Check jersey number uniqueness
		jerseyNumber := player.JerseyNumber
		if req.JerseyNumber > 0 {
			jerseyNumber = req.JerseyNumber
		}
		
		exists, err := s.playerRepo.ExistsByJerseyNumberInTeam(*req.TeamID, jerseyNumber, id)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, errors.New("jersey number already exists in this team")
		}
		
		player.TeamID = req.TeamID
	}

	if req.Name != "" {
		player.Name = req.Name
	}
	if req.Height > 0 {
		player.Height = req.Height
	}
	if req.Weight > 0 {
		player.Weight = req.Weight
	}
	if req.Position != "" {
		player.Position = domain.PlayerPosition(req.Position)
	}
	if req.JerseyNumber > 0 {
		player.JerseyNumber = req.JerseyNumber
	}

	if err := s.playerRepo.Update(player); err != nil {
		return nil, err
	}

	// Load team relation
	player, _ = s.playerRepo.FindByIDWithTeam(player.ID)
	return player, nil
}

func (s *playerService) Delete(id uint) error {
	player, err := s.playerRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("player not found")
		}
		return err
	}

	return s.playerRepo.Delete(player.ID)
}
