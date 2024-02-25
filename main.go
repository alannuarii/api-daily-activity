package main

import (
	// "fmt"

	"api-daily-activity/controllers"
	"api-daily-activity/db"
	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()

	db.ConnectionDatabase()

	r.GET("/api/activities", controllers.GetAllData)

	r.POST("/api/activity", controllers.PostActivity)

	r.Run(":8888")
}