package domain

import "time"

type Rover struct {
	ID        uint `gorm:"primaryKey"`
	Pos_x     int  `gorm:"not_null;default:0"`
	Pos_y     int  `gorm:"not_null;default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Rover) TableName() string {
	return "rovers"
}
