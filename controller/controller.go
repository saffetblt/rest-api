package controller

import (
	"rest-api/database"
	"rest-api/model"

	"github.com/gin-gonic/gin"
)

type User = model.User

func GetUsers(c *gin.Context) {
	db := database.GetDB()
	users := []User{}

	db.Find(&users)
	c.JSON(200, users)
}

func GetUser(c *gin.Context) {
	db := database.GetDB()
	id := c.Params.ByName("id")
	var user User

	if db.First(&user, id).Error != nil {
		c.AbortWithStatus(404)
	} else {
		c.JSON(200, user)
	}

}

func CreateUser(c *gin.Context) {
	db := database.GetDB()
	var user User

	c.BindJSON(&user)
	db.Create(&user)
	c.JSON(200, user)
}

func DeleteUser(c *gin.Context) {
	db := database.GetDB()
	id := c.Params.ByName("id")
	var user User

	if db.Delete(&user, id).Error != nil {
		c.AbortWithStatus(404)
	} else {
		c.JSON(200, gin.H{"id#" + id: "deleted"})
	}
}

func UpdateUser(c *gin.Context) {
	db := database.GetDB()
	id := c.Params.ByName("id")
	var user User
	if db.First(&user, id).Error != nil {
		c.AbortWithStatus(404)
	} else {
		c.BindJSON(&user)
		db.Save(&user)
		c.JSON(200, user)
	}
}
