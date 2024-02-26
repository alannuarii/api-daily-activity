package controllers

import (
	"api-daily-activity/db"
	"api-daily-activity/models"
	"fmt"
	"net/http"
	// "path/filepath"

	"github.com/gin-gonic/gin"
)


func GetAllData(c *gin.Context) {
	// Menggunakan koneksi DB yang sudah dibuat di file connection.go
	db := db.DB

	rows, err := db.Query("SELECT * FROM activity")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal mendapatkan data dari database"})
		return
	}
	// defer rows.Close()

	var activities []models.Activity

	for rows.Next() {
		var activity models.Activity
		if err := rows.Scan(&activity.ID, &activity.Tanggal, &activity.Jenis, &activity.Perusahaan, &activity.Pekerjaan, &activity.Kode); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal membaca data dari hasil query"})
			return
		}
		activities = append(activities, activity)
	}

	if err := rows.Err(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal membaca data dari hasil query"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sukses", "data": activities})
}



func PostActivity(c *gin.Context) {
	var activity models.Activity

	files := c.Request.MultipartForm.File["foto"]
	fmt.Println(files)
    // for _, file := range files {
    //     // Simpan file ke server
    //     err := c.SaveUploadedFile(file, filepath.Join("static/img", file.Filename))
    //     if err != nil {
    //         c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal menyimpan file"})
    //         fmt.Println("Error menyimpan file:", err)
    //         return
    //     }
    // }

	if err := c.ShouldBindJSON(&activity); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		fmt.Println("Error 1:", err)
		return
	}
	fmt.Println(activity)
	db := db.DB

	query := `INSERT INTO activity (tanggal, jenis, perusahaan, pekerjaan, kode) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	err := db.QueryRow(query, activity.Tanggal, activity.Jenis, activity.Perusahaan, activity.Pekerjaan, activity.Kode).Scan(&activity.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal membuat entri di database"})
		fmt.Println("Error 2:", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sukses", "data": activity})
}





