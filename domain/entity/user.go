package entity

import "time"

type User struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"uniqueIndex;size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
