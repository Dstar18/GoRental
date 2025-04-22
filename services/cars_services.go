package services

import (
	"GoRental/model"

	"gorm.io/gorm"
)

type CarsService struct {
	DB *gorm.DB
}

func NewCarsService(db *gorm.DB) *CarsService {
	return &CarsService{DB: db}
}

func (s *CarsService) GetsCars() ([]model.Cars, error) {
	var articles []model.Cars
	result := s.DB.Find(&articles).Error
	return articles, result
}

func (s *CarsService) GetIdCars(id uint) (model.Cars, error) {
	var car model.Cars
	err := s.DB.First(&car, id).Error
	return car, err
}

func (s *CarsService) CreateCars(data *model.Cars) error {
	return s.DB.Create(data).Error
}

func (s *CarsService) UpdateCars(data *model.Cars) error {
	return s.DB.Save(data).Error
}

func (s *CarsService) DeleteCars(id uint) error {
	return s.DB.Delete(&model.Cars{}, id).Error
}
