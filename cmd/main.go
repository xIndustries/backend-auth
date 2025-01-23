package main

import (
	"fmt"
	"log"
	"time"

	"github.com/fatih/color"

	"github.com/xIndustries/BandRoom/backend-auth/config"
	"github.com/xIndustries/BandRoom/backend-auth/db"
	"github.com/xIndustries/BandRoom/backend-auth/internal/handlers"
	"github.com/xIndustries/BandRoom/backend-auth/internal/repositories"
	"github.com/xIndustries/BandRoom/backend-auth/internal/server"
	"github.com/xIndustries/BandRoom/backend-auth/internal/services"
	"github.com/xIndustries/BandRoom/backend-auth/internal/utils"
)

func main() {
	showStartupBanner()

	// Load configuration
	cfg := config.LoadConfig()
	if cfg == nil {
		renderError("Failed to load configuration")
		log.Fatal("Failed to load configuration")
	}
	renderSuccess("Configuration loaded successfully")

	// Initialize logger
	err := utils.InitLogger("log/user-service.log")
	if err != nil {
		renderError(fmt.Sprintf("Failed to initialize logger: %v", err))
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	renderSuccess("Logger initialized successfully")

	// Connect to the database
	database, err := db.ConnectDB(cfg)
	if err != nil {
		renderError(fmt.Sprintf("Database connection failed: %v", err))
		log.Fatalf("Database connection failed: %v", err)
	}
	defer database.Close()
	renderSuccess("Connected to the database successfully")

	// Initialize repositories
	userRepo := repositories.NewUserRepository(database)
	renderStep("User repository initialized")

	// Initialize services
	userService := services.NewUserService(userRepo)
	renderStep("User service initialized")

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userService)
	renderStep("User handler initialized")

	// Start gRPC server
	serverPort := cfg.GRPCPort
	renderAction(fmt.Sprintf("Starting gRPC server on port %s", serverPort))
	if err := server.RunGRPCServer(serverPort, userHandler); err != nil {
		renderError(fmt.Sprintf("Failed to start gRPC server: %v", err))
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
	renderSuccess(fmt.Sprintf("gRPC server is listening on port %s", serverPort))
}

func showStartupBanner() {
	color.Cyan(`
==========================================================
                       X INDUSTRIES
==========================================================
`)
	time.Sleep(1 * time.Second)
	color.Green("Initializing System...")
	time.Sleep(500 * time.Millisecond)
}

func renderSuccess(message string) {
	color.Green("[SUCCESS] %s", message)
}

func renderError(message string) {
	color.Red("[ERROR] %s", message)
}

func renderStep(message string) {
	color.Cyan("[STEP] %s", message)
}

func renderAction(message string) {
	color.Yellow("[ACTION] %s", message)
	time.Sleep(200 * time.Millisecond) // Simulate a slight delay for effect
}
