package service

import (
	"crudproject/internal/models"
	"crudproject/internal/repository"
)

type CarService interface {
	GetAll() ([]models.Car, error)
	GetByID(id uint) (models.Car, error)
	Create(car models.Car) (models.Car, error)
	Update(car models.Car) (models.Car, error)
	Delete(id uint) error
}

type carService struct {
	repo repository.CarRepository
}

func NewCarService(repo repository.CarRepository) CarService {
	return &carService{repo}
}

func (s *carService) GetAll() ([]models.Car, error) {
	return s.repo.GetAll()
}

func (s *carService) GetByID(id uint) (models.Car, error) {
	return s.repo.GetByID(id)
}

func (s *carService) Create(car models.Car) (models.Car, error) {
	return s.repo.Create(car)
}

func (s *carService) Update(car models.Car) (models.Car, error) {
	return s.repo.Update(car)
}

func (s *carService) Delete(id uint) error {
	return s.repo.Delete(id)
}
