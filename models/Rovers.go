package models

import "gorm.io/gorm"

type Rover struct {
	gorm.Model

	id uint `gorm:"not null;unique_index"`
	pos_x int `gorm:"not_null;default:0"`
	pos_y int `gorm:"not_null;default:0"`
}
