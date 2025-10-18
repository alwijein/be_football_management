package main

import (
	"log"

	"github.com/alwijein/be_football/internal/config"
	"github.com/alwijein/be_football/internal/handler"
	"github.com/alwijein/be_football/internal/middleware"
	"github.com/alwijein/be_football/internal/repository"
	"github.com/alwijein/be_football/internal/routes"
	"github.com/alwijein/be_football/internal/service"
	"github.com/alwijein/be_football/pkg/database"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()
	log.Printf("Starting %s v%s\n", cfg.App.Name, cfg.App.Version)

	// Set Gin mode
	gin.SetMode(cfg.Server.GinMode)

	// Connect to database
	db, err := database.ConnectDatabase(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run migrations
	if err := database.AutoMigrate(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	teamRepo := repository.NewTeamRepository(db)
	playerRepo := repository.NewPlayerRepository(db)
	matchRepo := repository.NewMatchRepository(db)
	goalRepo := repository.NewGoalRepository(db)

	// Initialize services
	authService := service.NewAuthService(userRepo)
	teamService := service.NewTeamService(teamRepo)
	playerService := service.NewPlayerService(playerRepo, teamRepo)
	matchService := service.NewMatchService(matchRepo, teamRepo, playerRepo, goalRepo)
	reportService := service.NewReportService(matchRepo, goalRepo, teamRepo, playerRepo)

	// Initialize handlers
	authHandler := handler.NewAuthHandler(authService)
	teamHandler := handler.NewTeamHandler(teamService)
	playerHandler := handler.NewPlayerHandler(playerService)
	matchHandler := handler.NewMatchHandler(matchService)
	reportHandler := handler.NewReportHandler(reportService)

	// Initialize Gin router
	router := gin.Default()

	// Apply global middleware
	router.Use(middleware.CORS())
	router.Use(middleware.Logger())
	router.Use(middleware.ErrorHandler())

	// Setup routes
	routes.SetupRoutes(router, authHandler, teamHandler, playerHandler, matchHandler, reportHandler)

	// Start server on all network interfaces
	serverAddr := "0.0.0.0:" + cfg.Server.Port
	log.Printf("Server is running on:\n")
	log.Printf("  - Local:   http://localhost:%s\n", cfg.Server.Port)
	log.Printf("  - Network: http://<your-ip>:%s\n", cfg.Server.Port)
	log.Printf("  Listening on %s\n", serverAddr)

	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
