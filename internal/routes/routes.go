package routes

import (
	"crudproject/internal/auth"
	"crudproject/internal/handler"
	"crudproject/internal/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	h := handler.NewCarHandler(db)

	// Auth Routes
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", auth.Register)
		authRoutes.POST("/login", auth.Login)
		authRoutes.GET("/me", middleware.AuthRequired(), auth.Me)
	}

	// Car Routes (Protected)
	carRoutes := r.Group("/cars")
	carRoutes.Use(middleware.AuthRequired()) // 👈 тек токен арқылы кіреді
	{
		carRoutes.GET("/", h.GetAll)
		carRoutes.GET("/:id", h.GetByID)
		carRoutes.POST("/", h.Create)
		carRoutes.PUT("/:id", h.Update)
		carRoutes.DELETE("/:id", h.Delete)
	}
}
