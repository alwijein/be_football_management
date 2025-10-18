package service

import (
	"errors"

	"github.com/alwijein/be_football/internal/domain"
	"github.com/alwijein/be_football/internal/dto"
	"github.com/alwijein/be_football/internal/repository"
	"gorm.io/gorm"
)

// TeamService handles team business logic
type TeamService interface {
	Create(req *dto.CreateTeamRequest) (*domain.Team, error)
	GetByID(id uint) (*domain.Team, error)
	GetByIDWithPlayers(id uint) (*domain.Team, error)
	GetAll(pagination *dto.Pagination) ([]domain.Team, *dto.Pagination, error)
	GetRegistered(limit int) ([]domain.Team, error)
	Update(id uint, req *dto.UpdateTeamRequest) (*domain.Team, error)
	Delete(id uint) error
}

type teamService struct {
	teamRepo repository.TeamRepository
}

// NewTeamService creates new team service
func NewTeamService(teamRepo repository.TeamRepository) TeamService {
	return &teamService{
		teamRepo: teamRepo,
	}
}

func (s *teamService) Create(req *dto.CreateTeamRequest) (*domain.Team, error) {
	// Check if team name already exists
	exists, err := s.teamRepo.ExistsByName(req.Name, 0)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("team name already exists")
	}

	team := &domain.Team{
		Name:            req.Name,
		Logo:            req.Logo,
		EstablishedYear: req.EstablishedYear,
		Address:         req.Address,
		City:            req.City,
	}

	if err := s.teamRepo.Create(team); err != nil {
		return nil, err
	}

	return team, nil
}

func (s *teamService) GetByID(id uint) (*domain.Team, error) {
	team, err := s.teamRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("team not found")
		}
		return nil, err
	}
	return team, nil
}

func (s *teamService) GetByIDWithPlayers(id uint) (*domain.Team, error) {
	team, err := s.teamRepo.FindByIDWithPlayers(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("team not found")
		}
		return nil, err
	}
	return team, nil
}

func (s *teamService) GetAll(pagination *dto.Pagination) ([]domain.Team, *dto.Pagination, error) {
	teams, err := s.teamRepo.FindAll(pagination)
	if err != nil {
		return nil, nil, err
	}

	total, err := s.teamRepo.Count()
	if err != nil {
		return nil, nil, err
	}

	pagination.Total = total
	pagination.CalculateTotalPages()

	return teams, pagination, nil
}

func (s *teamService) Update(id uint, req *dto.UpdateTeamRequest) (*domain.Team, error) {
	team, err := s.teamRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("team not found")
		}
		return nil, err
	}

	// Check if new name already exists
	if req.Name != "" && req.Name != team.Name {
		exists, err := s.teamRepo.ExistsByName(req.Name, id)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, errors.New("team name already exists")
		}
		team.Name = req.Name
	}

	if req.Logo != "" {
		team.Logo = req.Logo
	}
	if req.EstablishedYear > 0 {
		team.EstablishedYear = req.EstablishedYear
	}
	if req.Address != "" {
		team.Address = req.Address
	}
	if req.City != "" {
		team.City = req.City
	}

	if err := s.teamRepo.Update(team); err != nil {
		return nil, err
	}

	return team, nil
}

func (s *teamService) Delete(id uint) error {
	team, err := s.teamRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("team not found")
		}
		return err
	}

	return s.teamRepo.Delete(team.ID)
}

func (s *teamService) GetRegistered(limit int) ([]domain.Team, error) {
	// Get teams with default limit of 5 for homepage
	if limit <= 0 {
		limit = 5
	}

	pagination := &dto.Pagination{
		Page:  1,
		Limit: limit,
	}

	teams, _, err := s.GetAll(pagination)
	if err != nil {
		return nil, err
	}

	return teams, nil
}
