package main

import (
	"api-daily-activity/controllers"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.GET("/api/activities", controllers.GetAllActivity)
    r.GET("/api/activity1/:kode", controllers.GetActivityParam1)
    r.GET("/api/activity2/:kode", controllers.GetActivityParam2)

    r.POST("/api/activity", controllers.PostActivity)
    r.POST("/api/photos/:kode", controllers.PostPhotos)

    r.Run(":8888")
}
