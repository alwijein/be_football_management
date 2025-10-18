package repository

import (
	"github.com/alwijein/be_football/internal/domain"
	"github.com/alwijein/be_football/internal/dto"
	"gorm.io/gorm"
)

// MatchRepository handles match data operations
type MatchRepository interface {
	Create(match *domain.Match) error
	FindByID(id uint) (*domain.Match, error)
	FindByIDWithRelations(id uint) (*domain.Match, error)
	FindAll(pagination *dto.Pagination) ([]domain.Match, error)
	FindByTeamID(teamID uint) ([]domain.Match, error)
	Count() (int64, error)
	Update(match *domain.Match) error
	Delete(id uint) error
	CountTeamWins(teamID uint, upToMatchID uint) (int64, error)
}

type matchRepository struct {
	db *gorm.DB
}

// NewMatchRepository creates new match repository
func NewMatchRepository(db *gorm.DB) MatchRepository {
	return &matchRepository{db: db}
}

func (r *matchRepository) Create(match *domain.Match) error {
	return r.db.Create(match).Error
}

func (r *matchRepository) FindByID(id uint) (*domain.Match, error) {
	var match domain.Match
	err := r.db.First(&match, id).Error
	if err != nil {
		return nil, err
	}
	return &match, nil
}

func (r *matchRepository) FindByIDWithRelations(id uint) (*domain.Match, error) {
	var match domain.Match
	err := r.db.Preload("HomeTeam").
		Preload("AwayTeam").
		Preload("Goals.Player.Team").
		First(&match, id).Error
	if err != nil {
		return nil, err
	}
	return &match, nil
}

func (r *matchRepository) FindAll(pagination *dto.Pagination) ([]domain.Match, error) {
	var matches []domain.Match
	err := r.db.Preload("HomeTeam").
		Preload("AwayTeam").
		Order("match_date DESC, match_time DESC").
		Limit(pagination.GetLimit()).
		Offset(pagination.GetOffset()).
		Find(&matches).Error
	return matches, err
}

func (r *matchRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&domain.Match{}).Count(&count).Error
	return count, err
}

func (r *matchRepository) Update(match *domain.Match) error {
	return r.db.Save(match).Error
}

func (r *matchRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Match{}, id).Error
}

// CountTeamWins counts total wins for a team up to a specific match
func (r *matchRepository) CountTeamWins(teamID uint, upToMatchID uint) (int64, error) {
	var count int64

	// Count as home team wins
	var homeWins int64
	err := r.db.Model(&domain.Match{}).
		Where("id <= ? AND home_team_id = ? AND status = ? AND home_score > away_score",
			upToMatchID, teamID, domain.StatusCompleted).
		Count(&homeWins).Error
	if err != nil {
		return 0, err
	}

	// Count as away team wins
	var awayWins int64
	err = r.db.Model(&domain.Match{}).
		Where("id <= ? AND away_team_id = ? AND status = ? AND away_score > home_score",
			upToMatchID, teamID, domain.StatusCompleted).
		Count(&awayWins).Error
	if err != nil {
		return 0, err
	}

	count = homeWins + awayWins
	return count, nil
}

func (r *matchRepository) FindByTeamID(teamID uint) ([]domain.Match, error) {
	var matches []domain.Match
	err := r.db.
		Preload("HomeTeam").
		Preload("AwayTeam").
		Preload("Goals.Player.Team").
		Where("home_team_id = ? OR away_team_id = ?", teamID, teamID).
		Order("match_date DESC, match_time DESC").
		Find(&matches).Error

	return matches, err
}
