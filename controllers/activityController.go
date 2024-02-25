package controllers

import (
	"api-daily-activity/db"
	"api-daily-activity/models"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
)


func GetAllData(c *gin.Context){
	var activities []models.Activity

	db.DB.Find(&activities)
	c.JSON(http.StatusOK, gin.H{"message":"Sukses", "data":activities})
}

func PostActivity(c *gin.Context) {
	
	var activity models.Activity

	// Membaca seluruh isi payload JSON dari permintaan HTTP
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		// Tangani kesalahan jika gagal membaca payload JSON
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal membaca payload JSON"})
		fmt.Println("Error:", err)
		return
	}

	// Menampilkan payload JSON yang diterima di terminal
	fmt.Println("Request JSON:", string(body))

	// Mengembalikan payload JSON ke dalam Request Body
	c.Request.Body = ioutil.NopCloser(bytes.NewReader(body))

	// Parsing tanggal yang diterima dari JSON
	tanggal, err := time.Parse("2006-01-02", activity.Tanggal.String())
	if err != nil {
		// Tangani kesalahan jika gagal mem-parse tanggal
		c.JSON(http.StatusBadRequest, gin.H{"message": "Gagal mem-parse tanggal", "error": err.Error()})
		fmt.Println("Error 2:", err)
		return
	}

	activity.Tanggal = tanggal

	// Menerima JSON dari permintaan HTTP dan mengikatnya ke struct Activity
	if err := c.ShouldBindJSON(&activity); err != nil {
		// Tangani kesalahan jika terjadi kesalahan dalam pengikatan JSON
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		fmt.Println("Error 1:", err)
		return
	}

	// Buat entri baru di database menggunakan struct Activity yang sudah diisi
	if err := db.DB.Create(&activity).Error; err != nil {
		// Tangani kesalahan jika terjadi kesalahan saat membuat entri di database
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal membuat entri di database"})
		fmt.Println("Error 3:", err)
		return
	}

	// Kirim respons ke klien untuk memberi tahu bahwa permintaan berhasil
	c.JSON(http.StatusOK, gin.H{"message": "Sukses", "data": activity})
}





