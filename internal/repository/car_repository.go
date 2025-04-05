package repository

import (
	"crudproject/internal/models"
	"gorm.io/gorm"
)

type CarRepository interface {
	GetAll() ([]models.Car, error)
	GetByID(id uint) (models.Car, error)
	Create(car models.Car) (models.Car, error)
	Update(car models.Car) (models.Car, error)
	Delete(id uint) error
}

type carRepo struct {
	db *gorm.DB
}

func NewCarRepository(db *gorm.DB) CarRepository {
	return &carRepo{db}
}

func (r *carRepo) GetAll() ([]models.Car, error) {
	var cars []models.Car
	err := r.db.Find(&cars).Error
	return cars, err
}

func (r *carRepo) GetByID(id uint) (models.Car, error) {
	var car models.Car
	err := r.db.First(&car, id).Error
	return car, err
}

func (r *carRepo) Create(car models.Car) (models.Car, error) {
	err := r.db.Create(&car).Error
	return car, err
}

func (r *carRepo) Update(car models.Car) (models.Car, error) {
	err := r.db.Save(&car).Error
	return car, err
}

func (r *carRepo) Delete(id uint) error {
	return r.db.Delete(&models.Car{}, id).Error
}
