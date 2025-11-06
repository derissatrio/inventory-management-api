package http

import (
	"github.com/gin-gonic/gin"
	applicationservice "inventory-ticketing-system/application/service"
	"inventory-ticketing-system/application/usecase/asset"
	"inventory-ticketing-system/application/usecase/auth"
	"inventory-ticketing-system/application/usecase/ticket"
	"inventory-ticketing-system/delivery/http/handler"
	"inventory-ticketing-system/delivery/http/middleware"
	"inventory-ticketing-system/infrastructure/jwt"
)

type Router struct {
	engine *gin.Engine
}

func NewRouter(
	authHandler *handler.AuthHandler,
	assetHandler *handler.AssetHandler,
	ticketHandler *handler.TicketHandler,
	locationHandler *handler.LocationHandler,
	jwtManager *jwt.JWTManager,
) *Router {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

	// Middleware
	engine.Use(middleware.LoggingMiddleware())
	engine.Use(middleware.CORSMiddleware())
	engine.Use(gin.Recovery())

	// Routes
	router := &Router{
		engine: engine,
	}

	router.setupRoutes(authHandler, assetHandler, ticketHandler, locationHandler, jwtManager)

	return router
}

func (r *Router) setupRoutes(
	authHandler *handler.AuthHandler,
	assetHandler *handler.AssetHandler,
	ticketHandler *handler.TicketHandler,
	locationHandler *handler.LocationHandler,
	jwtManager *jwt.JWTManager,
) {
	v1 := r.engine.Group("/api/v1")

	// Health check route (public)
	v1.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
			"message": "Inventory & Ticketing Management System is running",
		})
	})

	// Public routes
	authRoutes := v1.Group("/auth")
	{
		authRoutes.POST("/login", authHandler.Login)
	}

	// Protected routes
	protected := v1.Group("/")
	protected.Use(middleware.AuthMiddleware(jwtManager))
	{
		// Asset routes
		assetRoutes := protected.Group("/assets")
		{
			assetRoutes.GET("", assetHandler.List)                   // All authenticated users
			assetRoutes.GET("/:id", assetHandler.Get)               // All authenticated users
			assetRoutes.POST("", middleware.RoleMiddleware("admin"), assetHandler.Create) // Admin only
			assetRoutes.PUT("/:id", middleware.RoleMiddleware("admin"), assetHandler.Update) // Admin only
			assetRoutes.DELETE("/:id", middleware.RoleMiddleware("admin"), assetHandler.Delete) // Admin only
		}

		// Ticket routes
		ticketRoutes := protected.Group("/tickets")
		{
			ticketRoutes.GET("", ticketHandler.List)                 // All authenticated users
			ticketRoutes.GET("/:id", ticketHandler.Get)             // All authenticated users
			ticketRoutes.POST("", ticketHandler.Create)             // All authenticated users
			ticketRoutes.PUT("/:id", middleware.RoleMiddleware("admin"), ticketHandler.Update) // Admin only
			ticketRoutes.DELETE("/:id", middleware.RoleMiddleware("admin"), ticketHandler.Delete) // Admin only
		}

		// Location routes
		locationRoutes := protected.Group("/locations")
		{
			locationRoutes.GET("", locationHandler.List)                 // All authenticated users
			locationRoutes.GET("/:id", locationHandler.Get)             // All authenticated users
			locationRoutes.POST("", middleware.RoleMiddleware("admin"), locationHandler.Create) // Admin only
			locationRoutes.PUT("/:id", middleware.RoleMiddleware("admin"), locationHandler.Update) // Admin only
			locationRoutes.DELETE("/:id", middleware.RoleMiddleware("admin"), locationHandler.Delete) // Admin only
		}

		// User routes
		userRoutes := protected.Group("/users")
		{
			userRoutes.GET("/me", func(c *gin.Context) {
				userID, _ := c.Get("user_id")
				userRole, _ := c.Get("user_role")
				c.JSON(200, gin.H{
					"success": true,
					"data": gin.H{
						"id":   userID,
						"role": userRole,
					},
					"message": "User profile retrieved successfully",
				})
			})
		}
	}
}

func (r *Router) GetEngine() *gin.Engine {
	return r.engine
}

func SetupDependencies() (
	*handler.AuthHandler,
	*handler.AssetHandler,
	*handler.TicketHandler,
	*handler.LocationHandler,
	*jwt.JWTManager,
) {
	// Note: This function is deprecated. Use cmd/app/main.go for proper dependency injection
	// This is kept only for compatibility

	// JWT Manager
	jwtManager := jwt.NewJWTManager("your-secret-key")

	// Use Cases - these are mocked for now
	loginUseCase := auth.NewLoginUseCase(nil)
	createAssetUseCase := asset.NewCreateAssetUseCase(nil)
	listAssetsUseCase := asset.NewListAssetsUseCase(nil)
	createTicketUseCase := ticket.NewCreateTicketUseCase(nil)
	listTicketsUseCase := ticket.NewListTicketsUseCase(nil)

	// Handlers - location handler needs a service, pass nil for now (should be injected from main)
	locationService := applicationservice.NewLocationService(nil)
	locationHandler := handler.NewLocationHandler(locationService)

	authHandler := handler.NewAuthHandler(loginUseCase)
	assetHandler := handler.NewAssetHandler(createAssetUseCase, listAssetsUseCase)
	ticketHandler := handler.NewTicketHandler(createTicketUseCase, listTicketsUseCase)

	return authHandler, assetHandler, ticketHandler, locationHandler, jwtManager
}