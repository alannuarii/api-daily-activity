package controllers

import (
	"api-daily-activity/db"
	"path/filepath"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)


func PostActivity(c *gin.Context) {

    db := db.DB
    // Parsing data dari FormData
    err := c.Request.ParseMultipartForm(0)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal memproses data FormData"})
        fmt.Println("Error parsing FormData:", err)
        return
    }

    // Mengambil nilai-nilai dari FormData secara manual
    jenis := c.Request.FormValue("jenis")
    tanggal := c.Request.FormValue("tanggal")
    pekerjaan := c.Request.FormValue("pekerjaan")
    kode := c.Request.FormValue("kode")
    perusahaan := c.Request.FormValue("perusahaan")

    // Mengambil foto-foto dari FormData
    fotos := c.Request.MultipartForm.File["foto"]
    destination := "static/img"

    for index, foto := range fotos {

        rename := fmt.Sprintf("%s-%s-%s-%d.png", strings.ReplaceAll(pekerjaan, " ", "-"), tanggal, kode, index+1)

        query := `INSERT INTO photo (tanggal, kode, foto) VALUES ($1, $2, $3) RETURNING id`
        var activityID int
        err = db.QueryRow(query, tanggal, kode, rename).Scan(&activityID)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal membuat entri di database"})
            fmt.Println("Error menyimpan data ke database:", err)
            return
        }
        
        err := c.SaveUploadedFile(foto, filepath.Join(destination, rename))
        if err != nil {
            c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal menyimpan foto"})
            fmt.Println("Error menyimpan foto:", err)
            return
        }
    }

    query := `INSERT INTO activity (tanggal, jenis, perusahaan, pekerjaan, kode) VALUES ($1, $2, $3, $4, $5) RETURNING id`

    var activityID int
    err = db.QueryRow(query, tanggal, jenis, perusahaan, pekerjaan, kode).Scan(&activityID)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal membuat entri di database"})
        fmt.Println("Error menyimpan data ke database:", err)
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Sukses"})
}


func PostPhotos(c *gin.Context) {

    db := db.DB

    kodeData := c.Param("kode")

    err := c.Request.ParseMultipartForm(0)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal memproses data FormData"})
        fmt.Println("Error parsing FormData:", err)
        return
    }

    tanggal := c.Request.FormValue("tanggal")
    pekerjaan := c.Request.FormValue("pekerjaan")
    kode := kodeData

    fotos := c.Request.MultipartForm.File["foto"]
    destination := "static/img"

    for index, foto := range fotos {

        rename := fmt.Sprintf("%s-%s-%s-%d.png", strings.ReplaceAll(pekerjaan, " ", "-"), tanggal, kode, index+1)

        query := `INSERT INTO photo (tanggal, kode, foto) VALUES ($1, $2, $3) RETURNING id`
        var activityID int
        err = db.QueryRow(query, tanggal, kode, rename).Scan(&activityID)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal membuat entri di database"})
            fmt.Println("Error menyimpan data ke database:", err)
            return
        }
        
        err := c.SaveUploadedFile(foto, filepath.Join(destination, rename))
        if err != nil {
            c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal menyimpan foto"})
            fmt.Println("Error menyimpan foto:", err)
            return
        }
    }

    c.JSON(http.StatusOK, gin.H{"message": "Sukses"})
}