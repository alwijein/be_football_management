package repository

import (
	"github.com/alwijein/be_football/internal/domain"
	"github.com/alwijein/be_football/internal/dto"
	"gorm.io/gorm"
)

// GoalRepository handles goal data operations
type GoalRepository interface {
	Create(goal *domain.Goal) error
	FindByID(id uint) (*domain.Goal, error)
	FindByMatchID(matchID uint) ([]domain.Goal, error)
	Delete(id uint) error
	GetTopScorerInMatch(matchID uint) (*dto.TopScorerDTO, error)
	CountByPlayerInMatch(playerID uint, matchID uint) (int64, error)
}

type goalRepository struct {
	db *gorm.DB
}

// NewGoalRepository creates new goal repository
func NewGoalRepository(db *gorm.DB) GoalRepository {
	return &goalRepository{db: db}
}

func (r *goalRepository) Create(goal *domain.Goal) error {
	return r.db.Create(goal).Error
}

func (r *goalRepository) FindByID(id uint) (*domain.Goal, error) {
	var goal domain.Goal
	err := r.db.Preload("Player").Preload("Match").First(&goal, id).Error
	if err != nil {
		return nil, err
	}
	return &goal, nil
}

func (r *goalRepository) FindByMatchID(matchID uint) ([]domain.Goal, error) {
	var goals []domain.Goal
	err := r.db.Preload("Player.Team").
		Where("match_id = ?", matchID).
		Order("minute ASC").
		Find(&goals).Error
	return goals, err
}

func (r *goalRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Goal{}, id).Error
}

// GetTopScorerInMatch gets top scorer in a specific match
func (r *goalRepository) GetTopScorerInMatch(matchID uint) (*dto.TopScorerDTO, error) {
	var result dto.TopScorerDTO
	
	err := r.db.Model(&domain.Goal{}).
		Select("goals.player_id, players.name as player_name, teams.name as team_name, COUNT(*) as goal_count").
		Joins("JOIN players ON goals.player_id = players.id").
		Joins("LEFT JOIN teams ON players.team_id = teams.id").
		Where("goals.match_id = ?", matchID).
		Group("goals.player_id, players.name, teams.name").
		Order("goal_count DESC").
		Limit(1).
		Scan(&result).Error
		
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	
	if result.PlayerID == 0 {
		return nil, nil
	}
	
	return &result, nil
}

// CountByPlayerInMatch counts goals by player in a match
func (r *goalRepository) CountByPlayerInMatch(playerID uint, matchID uint) (int64, error) {
	var count int64
	err := r.db.Model(&domain.Goal{}).
		Where("player_id = ? AND match_id = ?", playerID, matchID).
		Count(&count).Error
	return count, err
}
