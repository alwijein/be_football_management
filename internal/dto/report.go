package dto

import "time"

// DashboardStatsResponse represents dashboard statistics
type DashboardStatsResponse struct {
	TotalTeams     int64 `json:"total_teams"`
	TotalSchedules int64 `json:"total_schedules"`
	TotalPlayers   int64 `json:"total_players"`
}

// MatchReportResponse represents detailed match report
type MatchReportResponse struct {
	ID           uint           `json:"id"`
	MatchDate    time.Time      `json:"match_date"`
	MatchTime    string         `json:"match_time"`
	HomeTeam     *TeamResponse  `json:"home_team"`
	AwayTeam     *TeamResponse  `json:"away_team"`
	HomeScore    *int           `json:"home_score"`
	AwayScore    *int           `json:"away_score"`
	MatchResult  string         `json:"match_result"`
	TopScorer    *TopScorerDTO  `json:"top_scorer"`
	HomeTeamWins int64          `json:"home_team_total_wins"`
	AwayTeamWins int64          `json:"away_team_total_wins"`
	Goals        []GoalResponse `json:"goals"`
}

// TopScorerDTO represents top scorer information
type TopScorerDTO struct {
	PlayerID   uint   `json:"player_id"`
	PlayerName string `json:"player_name"`
	TeamName   string `json:"team_name"`
	GoalCount  int64  `json:"goal_count"`
}

// TeamStatisticsResponse represents team statistics
type TeamStatisticsResponse struct {
	TeamID        uint   `json:"team_id"`
	TeamName      string `json:"team_name"`
	TotalMatches  int64  `json:"total_matches"`
	TotalWins     int64  `json:"total_wins"`
	TotalDraws    int64  `json:"total_draws"`
	TotalLosses   int64  `json:"total_losses"`
	GoalsScored   int64  `json:"goals_scored"`
	GoalsConceded int64  `json:"goals_conceded"`
}
