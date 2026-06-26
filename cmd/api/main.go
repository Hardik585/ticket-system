package main

import (
	"net/http"
	"strings"
	"ticket-system/internal/auth"
	"ticket-system/internal/repository"
	"ticket-system/internal/ticket"
	"github.com/gin-gonic/gin"
)

func main() {

	store := repository.NewMemoryStore()

	jwtSecret := "super_secret_signing_key_change_me_in_prod"
	authService := auth.NewService(jwtSecret, store)
	ticketService := ticket.NewService(store)

	authHandler := auth.NewHandler(authService)
	ticketHandler := ticket.NewHandler(ticketService)

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.POST("/auth/register", authHandler.Register)
	r.POST("/auth/login", authHandler.Login)

	protected := r.Group("/")
	protected.Use(AuthMiddleware(authService))
	{
		protected.POST("/tickets", ticketHandler.Create)
		protected.GET("/tickets", ticketHandler.List)
		protected.GET("/tickets/:id", ticketHandler.GetByID)
		protected.PATCH("/tickets/:id/status", ticketHandler.UpdateStatus)
	}
	r.Run(":8080")
}

func AuthMiddleware(authService *auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort() 
			return
		}
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be 'Bearer <token>'"})
			c.Abort()
			return
		}
		tokenString := parts[1]
		userID, err := authService.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Set("user_id", userID)
		c.Next()
	}
}