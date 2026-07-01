package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/customer-comm-dashboard/backend/internal/config"
	"github.com/customer-comm-dashboard/backend/internal/database"
	"github.com/customer-comm-dashboard/backend/internal/middleware"
	"github.com/customer-comm-dashboard/backend/internal/modules/auth"
	"github.com/customer-comm-dashboard/backend/internal/modules/inbox"
	"github.com/customer-comm-dashboard/backend/internal/response"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Connect to PostgreSQL
	if err := database.ConnectPostgres(cfg.DatabaseURL); err != nil {
		log.Printf("WARNING: PostgreSQL connection failed: %v", err)
	} else {
		defer database.ClosePostgres()

		// Run migrations
		authRepo := auth.NewRepository(database.Pool)
		migCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := authRepo.RunMigrations(migCtx); err != nil {
			log.Printf("WARNING: Failed to run auth migrations: %v", err)
		} else {
			log.Println("Auth migrations completed successfully")
		}

		// Run inbox migrations
		inboxRepo := inbox.NewRepository(database.Pool)
		if err := inboxRepo.RunMigrations(migCtx); err != nil {
			log.Printf("WARNING: Failed to run inbox migrations: %v", err)
		} else {
			log.Println("Inbox migrations completed successfully")
		}
	}

	// Connect to Redis
	if err := database.ConnectRedis(cfg.RedisAddr, cfg.RedisPass); err != nil {
		log.Printf("WARNING: Redis connection failed: %v", err)
	} else {
		defer database.CloseRedis()
	}

	// Connect to MinIO
	if err := database.ConnectMinio(cfg.MinioEndpoint, cfg.MinioAccessKey, cfg.MinioSecretKey, cfg.MinioBucket, cfg.MinioUseSSL); err != nil {
		log.Printf("WARNING: MinIO connection failed: %v", err)
	}

	// Initialize auth module
	authRepo := auth.NewRepository(database.Pool)
	authService := auth.NewService(authRepo, cfg.JWTSecret)
	authHandler := auth.NewHandler(authService)

	// Initialize inbox module
	inboxRepo := inbox.NewRepository(database.Pool)
	inboxService := inbox.NewService(inboxRepo)
	inboxHandler := inbox.NewHandler(inboxService)

	// Setup Gin router
	r := gin.Default()

	// Health check endpoint (public)
	r.GET("/health", func(c *gin.Context) {
		appStatus := "ok"
		postgresStatus := "ok"
		redisStatus := "ok"
		minioStatus := "ok"

		// Check PostgreSQL
		if database.Pool != nil {
			if err := database.Pool.Ping(c.Request.Context()); err != nil {
				postgresStatus = "error"
				log.Printf("Health check: PostgreSQL ping failed: %v", err)
			}
		} else {
			postgresStatus = "error"
		}

		// Check Redis
		if database.RedisClient != nil {
			if err := database.RedisClient.Ping(c.Request.Context()).Err(); err != nil {
				redisStatus = "error"
				log.Printf("Health check: Redis ping failed: %v", err)
			}
		} else {
			redisStatus = "error"
		}

		// Check MinIO
		if database.MinioClient != nil {
			if _, err := database.MinioClient.ListBuckets(c.Request.Context()); err != nil {
				minioStatus = "error"
				log.Printf("Health check: MinIO list buckets failed: %v", err)
			}
		} else {
			minioStatus = "error"
		}

		c.JSON(http.StatusOK, response.APIResponse{
			Success: true,
			Message: "Service is healthy",
			Data: gin.H{
				"app":      appStatus,
				"postgres": postgresStatus,
				"redis":    redisStatus,
				"minio":    minioStatus,
			},
		})
	})

	// Public auth routes
	authGroup := r.Group("/api/auth")
	{
		authGroup.POST("/login", authHandler.Login)
	}

	// Protected auth routes
	protectedAuth := r.Group("/api/auth")
	protectedAuth.Use(middleware.AuthRequired(authService))
	{
		protectedAuth.GET("/me", authHandler.Me)
	}

	// Protected inbox routes
	inboxGroup := r.Group("/api/inbox")
	inboxGroup.Use(middleware.AuthRequired(authService))
	{
		inboxGroup.GET("/conversations", inboxHandler.ListConversations)
		inboxGroup.GET("/conversations/:id", inboxHandler.GetConversation)
		inboxGroup.GET("/conversations/:id/messages", inboxHandler.ListMessages)
		inboxGroup.POST("/conversations/:id/messages", inboxHandler.SendMessage)
		inboxGroup.GET("/customers", inboxHandler.ListCustomers)
		inboxGroup.GET("/customers/:id", inboxHandler.GetCustomer)
	}

	// Start server
	addr := ":" + cfg.ServerPort
	log.Printf("Server starting on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
