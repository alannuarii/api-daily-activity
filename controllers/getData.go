package controllers

import (
	"api-daily-activity/db"
	"api-daily-activity/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetAllActivity(c *gin.Context) {

	db := db.DB
    // Query SQL menggunakan sqlx
    activities := []models.Activity{}
    err := db.Select(&activities, "SELECT * FROM activity")
    if err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal melakukan query ke database"})
        return
    }

    // Kirim respons dengan data kegiatan
    c.JSON(http.StatusOK, gin.H{"message": "Sukses", "data": activities})
}


// func GetActivityParam(c *gin.Context){

//     db := db.DB
// }