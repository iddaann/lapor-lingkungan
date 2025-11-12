package controllers

import (
	"laporan-lingkungan/config"
	"laporan-lingkungan/models"
	"laporan-lingkungan/middleware"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Create laporan
func CreateReport(c *gin.Context) {
	userID := middleware.GetUserID(c)

	var input models.Report
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	report := models.Report{
		Judul:        input.Judul,
		Deskripsi:    input.Deskripsi,
		JenisProblem: input.JenisProblem,
		Alamat:       input.Alamat,
		Latitude:     input.Latitude,
		Longitude:    input.Longitude,
		UserID:       userID,
		CreatedAt:    time.Now(),
	}

	config.DB.Create(&report)
	c.JSON(http.StatusOK, gin.H{"message": "Laporan berhasil dibuat", "report": report})
}

// Get semua laporan
func GetAllReports(c *gin.Context) {
	var reports []models.Report
	config.DB.Preload("User").Find(&reports)
	c.JSON(http.StatusOK, gin.H{"reports": reports})
}

// Get laporan milik user
func GetMyReports(c *gin.Context) {
	userID := middleware.GetUserID(c)
	var reports []models.Report
	config.DB.Where("user_id = ?", userID).Preload("User").Find(&reports)
	c.JSON(http.StatusOK, gin.H{"reports": reports})
}

// Update laporan
func UpdateReport(c *gin.Context) {
	id := c.Param("id")
	var report models.Report

	if err := config.DB.First(&report, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Laporan tidak ditemukan"})
		return
	}

	var input models.Report
	c.ShouldBindJSON(&input)

	report.Judul = input.Judul
	report.Deskripsi = input.Deskripsi
	report.JenisProblem = input.JenisProblem
	report.Alamat = input.Alamat
	report.Latitude = input.Latitude
	report.Longitude = input.Longitude
	config.DB.Save(&report)

	c.JSON(http.StatusOK, gin.H{"message": "Laporan berhasil diperbarui", "report": report})
}

// Delete laporan
func DeleteReport(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.Report{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Laporan berhasil dihapus"})
}

// Statistik cepat
func GetReportStats(c *gin.Context) {
	var total, today, pending, completed int64

	config.DB.Model(&models.Report{}).Count(&total)
	config.DB.Model(&models.Report{}).Where("DATE(created_at) = CURDATE()").Count(&today)
	config.DB.Model(&models.Report{}).Where("status = ?", "dilaporkan").Count(&pending)
	config.DB.Model(&models.Report{}).Where("status = ?", "selesai").Count(&completed)

	c.JSON(http.StatusOK, gin.H{
		"stats": gin.H{
			"total_reports":     total,
			"reports_today":     today,
			"pending_reports":   pending,
			"completed_reports": completed,
		},
	})
}
