package main

import (
	"fmt"
	"log"

	"inventory-ticketing-system/application/repository"
	"inventory-ticketing-system/application/service"
	"inventory-ticketing-system/application/usecase/asset"
	"inventory-ticketing-system/application/usecase/auth"
	"inventory-ticketing-system/application/usecase/ticket"
	httpdelivery "inventory-ticketing-system/delivery/http"
	"inventory-ticketing-system/delivery/http/handler"
	"inventory-ticketing-system/infrastructure/config"
	"inventory-ticketing-system/infrastructure/jwt"
	"inventory-ticketing-system/pkg/database"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize database
	db, err := database.NewDatabase(cfg.GetDatabaseDSN())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.CloseConnection(db)

	log.Println("Database connected successfully")

	// Initialize JWT manager
	jwtManager := jwt.NewJWTManager(cfg.JWTSecret)

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	assetRepo := repository.NewAssetRepository(db)
	ticketRepo := repository.NewTicketRepository(db)
	locationRepo := repository.NewLocationRepository(db)

	// Initialize services
	authService := service.NewAuthService(userRepo, jwtManager)
	assetService := service.NewAssetService(assetRepo)
	ticketService := service.NewTicketService(ticketRepo, assetRepo)
	locationService := service.NewLocationService(locationRepo)

	// Initialize use cases
	loginUseCase := auth.NewLoginUseCase(authService)
	createAssetUseCase := asset.NewCreateAssetUseCase(assetService)
	listAssetsUseCase := asset.NewListAssetsUseCase(assetService)
	createTicketUseCase := ticket.NewCreateTicketUseCase(ticketService)
	listTicketsUseCase := ticket.NewListTicketsUseCase(ticketService)

	// Initialize handlers
	authHandler := handler.NewAuthHandler(loginUseCase)
	assetHandler := handler.NewAssetHandler(createAssetUseCase, listAssetsUseCase)
	ticketHandler := handler.NewTicketHandler(createTicketUseCase, listTicketsUseCase)
	locationHandler := handler.NewLocationHandler(locationService)

	// Initialize router
	router := httpdelivery.NewRouter(authHandler, assetHandler, ticketHandler, locationHandler, jwtManager)

	// Start server
	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("Server starting on %s", addr)

	if err := router.GetEngine().Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
