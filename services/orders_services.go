package services

import (
	"GoRental/model"

	"gorm.io/gorm"
)

type OrdersService struct {
	DB *gorm.DB
}

func NewOrdersService(db *gorm.DB) *OrdersService {
	return &OrdersService{DB: db}
}

func (s *OrdersService) GetsOrders() ([]model.Orders, error) {
	var orders []model.Orders
	result := s.DB.Find(&orders).Error
	return orders, result
}

func (s *OrdersService) GetIdOrders(id uint) (model.Orders, error) {
	var order model.Orders
	err := s.DB.First(&order, id).Error
	return order, err
}
