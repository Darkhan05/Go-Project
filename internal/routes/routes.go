package routes

import (
	"crudproject/internal/auth"
	"crudproject/internal/handler"
	"crudproject/internal/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	authHandler := auth.NewAuthHandler(db)
	carHandler := handler.NewCarHandler(db)

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", authHandler.Register)
		authRoutes.POST("/login", authHandler.Login)
	}

	carRoutes := r.Group("/cars")
	carRoutes.Use(middleware.AuthRequired())
	{
		carRoutes.GET("/", carHandler.GetAll)
		carRoutes.GET("/:id", carHandler.GetByID)

		carRoutes.POST("/", middleware.AdminOnly(), carHandler.Create)
		carRoutes.PUT("/:id", middleware.AdminOnly(), carHandler.Update)
		carRoutes.DELETE("/:id", middleware.AdminOnly(), carHandler.Delete)
	}
}
