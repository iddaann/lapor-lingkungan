package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Nama      string    `json:"nama"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"-"`
	Role      string    `json:"role" gorm:"default:user"`
	CreatedAt time.Time `json:"created_at"`
}
