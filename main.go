package main

import (
	"crudproject/internal/handler"
	"crudproject/internal/models"
	"crudproject/internal/repository"
	"crudproject/internal/routes"
	"crudproject/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "postgres://postgres:7982@postgres:5432/postgres?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection error:", err)
	}

	if err := db.AutoMigrate(&models.Car{}); err != nil {
		log.Fatal("Migration error:", err)
	}

	carRepo := repository.NewCarRepository(db)
	carService := service.NewCarService(carRepo)
	carHandler := handler.NewCarHandler(carService)

	r := gin.Default()
	routes.SetupRoutes(r, carHandler)
	r.Run(":8080")
}
