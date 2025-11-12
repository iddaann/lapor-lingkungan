package controllers

import (
	"net/http"
	"time"

	"laporan-lingkungan/config"
	"laporan-lingkungan/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("secret_key")

func Register(c *gin.Context) {
	var input struct {
		Nama     string `json:"nama"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	user := models.User{Nama: input.Nama, Email: input.Email, Password: string(hash)}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email sudah digunakan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pendaftaran berhasil"})
}

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email tidak ditemukan"})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password salah"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, _ := token.SignedString(jwtKey)

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
		"user":  user,
	})
}
