package controllers

import (
	"api-daily-activity/db"
	"api-daily-activity/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetAllActivity(c *gin.Context) {

	db := db.DB

    activities := []models.Activity{}
    query := "SELECT * FROM activity"
    err := db.Select(&activities, query)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal melakukan query ke database"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Sukses", "data": activities})
}


func GetActivityParam1(c *gin.Context){

    db := db.DB
    kode := c.Param("kode")

    activity := []models.Activity{}
    query := "SELECT a.tanggal, a.jenis, a.pekerjaan, a.kode, a.perusahaan, p.foto FROM activity a JOIN photo p ON a.kode = p.kode WHERE a.kode = $1"
    err := db.Select(&activity, query, kode)

    if err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal melakukan query ke database"})
        fmt.Println("Error menyimpan data ke database:", err)
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Sukses", "data": activity})
}


func GetActivityParam2(c *gin.Context){

    db := db.DB
    kode := c.Param("kode")

    activity := []models.Activity{}
    query := "SELECT tanggal, jenis, pekerjaan, kode, perusahaan FROM activity WHERE kode = $1"
    err := db.Select(&activity, query, kode)

    if err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal melakukan query ke database"})
        fmt.Println("Error menyimpan data ke database:", err)
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Sukses", "data": activity})
}