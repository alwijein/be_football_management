package routes

import (
	"github.com/alwijein/be_football/internal/handler"
	"github.com/alwijein/be_football/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	router *gin.Engine,
	authHandler *handler.AuthHandler,
	teamHandler *handler.TeamHandler,
	playerHandler *handler.PlayerHandler,
	matchHandler *handler.MatchHandler,
	reportHandler *handler.ReportHandler,
) {
	// Serve static files (uploaded images)
	router.Static("/uploads", "./uploads")

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Football Management API is running",
		})
	})

	api := router.Group("/api")
	{
		api.POST("/login", authHandler.Login)

		protected := api.Group("")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.POST("/logout", authHandler.Logout)
			protected.GET("/dashboard/stats", reportHandler.GetDashboardStats)
			protected.GET("/matches/today", matchHandler.GetTodayMatches)

			protected.GET("/teams", teamHandler.GetAll)
			protected.GET("/teams/registered", teamHandler.GetRegistered)
			protected.POST("/teams", teamHandler.Create)
			protected.GET("/teams/:id/players", playerHandler.GetByTeamID)
			protected.POST("/teams/:id/players", playerHandler.Create)
			protected.GET("/teams/:id/match-reports", reportHandler.GetTeamMatchReports)
			protected.GET("/teams/:id", teamHandler.GetByID)
			protected.PUT("/teams/:id", teamHandler.Update)
			protected.DELETE("/teams/:id", teamHandler.Delete)

			protected.PUT("/players/:id", playerHandler.Update)
			protected.DELETE("/players/:id", playerHandler.Delete)

			protected.GET("/schedules", matchHandler.GetAll)
			protected.GET("/schedules/:id", matchHandler.GetByID)
			protected.POST("/schedules", matchHandler.Create)
			protected.PUT("/schedules/:id/result", matchHandler.SetResult)
			protected.DELETE("/schedules/:id", matchHandler.Delete)
			protected.GET("/schedules/:id/scorers", matchHandler.GetScorers)
			protected.POST("/schedules/:id/scorers", matchHandler.AddGoal)
			protected.DELETE("/scorers/:id", matchHandler.DeleteGoal)

			protected.GET("/profile", authHandler.GetProfile)
			protected.PUT("/profile", authHandler.UpdateProfile)
			protected.PUT("/profile/password", authHandler.ChangePassword)
		}
	}
}
