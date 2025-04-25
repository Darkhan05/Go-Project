package routes

import (
	"crudproject/internal/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	h := handler.NewCarHandler(db)

	r.GET("/cars", h.GetAll)
	r.GET("/cars/:id", h.GetByID)
	r.POST("/cars", h.Create)
	r.PUT("/cars/:id", h.Update)
	r.DELETE("/cars/:id", h.Delete)
}
