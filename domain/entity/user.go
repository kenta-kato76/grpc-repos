package entity

import "time"

type User struct {
	Name      string `gorm:"primaryKey"`
	Email     string `gorm:"uniqueIndex;size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
