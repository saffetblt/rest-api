package main

import (
	"rest-api/controller"
	"rest-api/database"

	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	port := "8085"
	url := "192.168.0.12"
	database.InitDB()

	r := gin.Default()
	user := r.Group("/user")
	{
		user.GET("/", controller.GetUsers)
		user.GET("/:id", controller.GetUser)
		user.POST("/", controller.CreateUser)
		user.DELETE("/:id", controller.DeleteUser)
		user.PUT("/:id", controller.UpdateUser)
	}
	r.Run(url + ":" + port)
}
