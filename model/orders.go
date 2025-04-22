package model

import "time"

type Orders struct {
	OrderID         uint      `gorm:"primaryKey" json:"order_id"`
	IDCar           int       `gorm:"type:bigint;not null" json:"id_car"` //foregn key ID Car
	OrderDate       time.Time `gorm:"type:timestamp;not null" json:"order_date"`
	PickupDate      time.Time `gorm:"type:timestamp;not null" json:"pickup_date"`
	DropOffDate     time.Time `gorm:"type:timestamp;not null" json:"dropoff_date"`
	PickupLocation  string    `gorm:"type:varchar(50);not null" json:"picup_location"`
	DropOffLocation string    `gorm:"type:varchar(50);not null" json:"dropoff_location"`
}
