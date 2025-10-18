package repository

import (
	"github.com/alwijein/be_football/internal/domain"
	"github.com/alwijein/be_football/internal/dto"
	"gorm.io/gorm"
)

// TeamRepository handles team data operations
type TeamRepository interface {
	Create(team *domain.Team) error
	FindByID(id uint) (*domain.Team, error)
	FindByIDWithPlayers(id uint) (*domain.Team, error)
	FindAll(pagination *dto.Pagination) ([]domain.Team, error)
	Count() (int64, error)
	Update(team *domain.Team) error
	Delete(id uint) error
	ExistsByName(name string, excludeID uint) (bool, error)
}

type teamRepository struct {
	db *gorm.DB
}

// NewTeamRepository creates new team repository
func NewTeamRepository(db *gorm.DB) TeamRepository {
	return &teamRepository{db: db}
}

func (r *teamRepository) Create(team *domain.Team) error {
	return r.db.Create(team).Error
}

func (r *teamRepository) FindByID(id uint) (*domain.Team, error) {
	var team domain.Team
	err := r.db.First(&team, id).Error
	if err != nil {
		return nil, err
	}
	return &team, nil
}

func (r *teamRepository) FindByIDWithPlayers(id uint) (*domain.Team, error) {
	var team domain.Team
	err := r.db.Preload("Players").First(&team, id).Error
	if err != nil {
		return nil, err
	}
	return &team, nil
}

func (r *teamRepository) FindAll(pagination *dto.Pagination) ([]domain.Team, error) {
	var teams []domain.Team
	err := r.db.Limit(pagination.GetLimit()).
		Offset(pagination.GetOffset()).
		Find(&teams).Error
	return teams, err
}

func (r *teamRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&domain.Team{}).Count(&count).Error
	return count, err
}

func (r *teamRepository) Update(team *domain.Team) error {
	return r.db.Save(team).Error
}

func (r *teamRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Team{}, id).Error
}

func (r *teamRepository) ExistsByName(name string, excludeID uint) (bool, error) {
	var count int64
	query := r.db.Model(&domain.Team{}).Where("name = ?", name)
	if excludeID > 0 {
		query = query.Where("id != ?", excludeID)
	}
	err := query.Count(&count).Error
	return count > 0, err
}
