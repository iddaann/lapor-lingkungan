package models

import "time"

type Report struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Judul        string    `json:"judul"`
	Deskripsi    string    `json:"deskripsi"`
	JenisProblem string    `json:"jenis_problem"`
	Alamat       string    `json:"alamat"`
	Latitude     float64   `json:"latitude"`
	Longitude    float64   `json:"longitude"`
	Status       string    `json:"status" gorm:"default:dilaporkan"`
	Prioritas    string    `json:"prioritas" gorm:"default:medium"`
	UserID       uint      `json:"user_id"`
	User         User      `json:"user" gorm:"foreignKey:UserID"`
	CreatedAt    time.Time `json:"created_at"`
}
