package main

import (
	"api-daily-activity/controllers"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.GET("/api/activities", controllers.GetAllActivity)
    r.POST("/api/activity", controllers.PostActivity)

    r.Run(":8888")
}
