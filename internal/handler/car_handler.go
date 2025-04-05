package handler

import (
	"crudproject/internal/models"
	"crudproject/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CarHandler struct {
	service service.CarService
}

func NewCarHandler(service service.CarService) *CarHandler {
	return &CarHandler{service}
}

func (h *CarHandler) GetAll(c *gin.Context) {
	cars, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cars"})
		return
	}
	c.JSON(http.StatusOK, cars)
}

func (h *CarHandler) GetByID(c *gin.Context) {
	// c.Param("id") мәнін uint типіне ауыстырып алу
	id := parseID(c.Param("id"))
	car, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
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
	created, err := h.service.Create(car)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create car"})
		return
	}
	c.JSON(http.StatusCreated, created)
}

func (h *CarHandler) Update(c *gin.Context) {
	var car models.Car
	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Update by ID
	car.ID = parseID(c.Param("id"))
	updated, err := h.service.Update(car)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update car"})
		return
	}
	c.JSON(http.StatusOK, updated)
}

func (h *CarHandler) Delete(c *gin.Context) {
	// c.Param("id") мәнін uint типіне ауыстырып алу
	id := parseID(c.Param("id"))
	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete car"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Car deleted"})
}

// parseID функциясы: id мәнін string-тен uint-ке ауыстырады
func parseID(id string) uint {
	var parsedID uint
	fmt.Sscanf(id, "%d", &parsedID)
	return parsedID
}
