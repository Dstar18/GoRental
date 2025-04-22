package model

type Cars struct {
	CarID     uint     `gorm:"primaryKey;not null" json:"car_id"`
	CarName   string   `gorm:"type:varchar(50);not null" json:"car_name"`
	DayRate   string   `gorm:"type:double;not null" json:"day_rate"`
	MonthRate string   `gorm:"type:double;not null" json:"month_rate"`
	Image     string   `gorm:"type:varchar(256);not null" json:"image"`
	Orders    []Orders `gorm:"foreignKey:IDCar" json:"orders"`
}
