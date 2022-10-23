package main

import (
	"github.com/colincheng/servicedesk/controllers"
	_ "github.com/colincheng/servicedesk/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/login", controllers.UserControl.Login)
	r.POST("/creatAdmin", controllers.UserControl.SaveUser)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
