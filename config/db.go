package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
    	os.Getenv("DB_USER"),
    	os.Getenv("DB_PASS"),
    	os.Getenv("DB_HOST"),
    	os.Getenv("DB_PORT"),
    	os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal konek ke database:", err)
	}

	DB = db
	fmt.Println("✅ Koneksi database berhasil")
}
