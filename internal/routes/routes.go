package routes

import (
	"crudproject/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, handler *handler.CarHandler) {
	r.GET("/cars", handler.GetAll)
	r.GET("/cars/:id", handler.GetByID)
	r.POST("/cars", handler.Create)
	r.PUT("/cars/:id", handler.Update)
	r.DELETE("/cars/:id", handler.Delete)
}
