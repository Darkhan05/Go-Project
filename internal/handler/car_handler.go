package handler

import (
	"crudproject/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type CarHandler struct {
	DB *gorm.DB
}

func NewCarHandler(db *gorm.DB) *CarHandler {
	return &CarHandler{DB: db}
}

// Мысал CRUD әдістері:
func (h *CarHandler) GetAll(c *gin.Context) {
	var cars []models.Car
	if err := h.DB.Find(&cars).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch cars"})
		return
	}
	c.JSON(http.StatusOK, cars)
}

func (h *CarHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var car models.Car
	if err := h.DB.First(&car, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "car not found"})
		return
	}
	c.JSON(http.StatusOK, car)
}

func (h *CarHandler) Create(c *gin.Context) {
	var car models.Car
	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.DB.Create(&car).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create car"})
		return
	}
	c.JSON(http.StatusCreated, car)
}

func (h *CarHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var car models.Car
	if err := h.DB.First(&car, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "car not found"})
		return
	}
	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.DB.Save(&car)
	c.JSON(http.StatusOK, car)
}

func (h *CarHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.DB.Delete(&models.Car{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete car"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "car deleted"})
}
