package controllers

import (
	"github.com/VeeramaniRajkumar32/ginLearn/initializers"
	"github.com/VeeramaniRajkumar32/ginLearn/models"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.Conn()
}

func PostCreate(c *gin.Context) {

	//create
	Post := models.Post{Title: "first", Body: "hello"}

	result := initializers.DB.Create(&Post)

	if result.Error != nil {
		c.Status(400)
	}
	c.JSON(200, gin.H{
		"post": Post,
	})
}
