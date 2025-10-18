package repository

import (
	"github.com/alwijein/be_football/internal/domain"
	"github.com/alwijein/be_football/internal/dto"
	"gorm.io/gorm"
)

// PlayerRepository handles player data operations
type PlayerRepository interface {
	Create(player *domain.Player) error
	FindByID(id uint) (*domain.Player, error)
	FindByIDWithTeam(id uint) (*domain.Player, error)
	FindAll(pagination *dto.Pagination) ([]domain.Player, error)
	FindByTeamID(teamID uint) ([]domain.Player, error)
	Count() (int64, error)
	Update(player *domain.Player) error
	Delete(id uint) error
	ExistsByJerseyNumberInTeam(teamID uint, jerseyNumber int, excludeID uint) (bool, error)
}

type playerRepository struct {
	db *gorm.DB
}

// NewPlayerRepository creates new player repository
func NewPlayerRepository(db *gorm.DB) PlayerRepository {
	return &playerRepository{db: db}
}

func (r *playerRepository) Create(player *domain.Player) error {
	return r.db.Create(player).Error
}

func (r *playerRepository) FindByID(id uint) (*domain.Player, error) {
	var player domain.Player
	err := r.db.First(&player, id).Error
	if err != nil {
		return nil, err
	}
	return &player, nil
}

func (r *playerRepository) FindByIDWithTeam(id uint) (*domain.Player, error) {
	var player domain.Player
	err := r.db.Preload("Team").First(&player, id).Error
	if err != nil {
		return nil, err
	}
	return &player, nil
}

func (r *playerRepository) FindAll(pagination *dto.Pagination) ([]domain.Player, error) {
	var players []domain.Player
	err := r.db.Preload("Team").
		Limit(pagination.GetLimit()).
		Offset(pagination.GetOffset()).
		Find(&players).Error
	return players, err
}

func (r *playerRepository) FindByTeamID(teamID uint) ([]domain.Player, error) {
	var players []domain.Player
	err := r.db.Where("team_id = ?", teamID).Find(&players).Error
	return players, err
}

func (r *playerRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&domain.Player{}).Count(&count).Error
	return count, err
}

func (r *playerRepository) Update(player *domain.Player) error {
	return r.db.Save(player).Error
}

func (r *playerRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Player{}, id).Error
}

func (r *playerRepository) ExistsByJerseyNumberInTeam(teamID uint, jerseyNumber int, excludeID uint) (bool, error) {
	var count int64
	query := r.db.Model(&domain.Player{}).
		Where("team_id = ? AND jersey_number = ?", teamID, jerseyNumber)
	if excludeID > 0 {
		query = query.Where("id != ?", excludeID)
	}
	err := query.Count(&count).Error
	return count > 0, err
}
