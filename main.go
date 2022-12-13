package main

import (
	"github.com/VeeramaniRajkumar32/ginLearn/controllers"
	"github.com/VeeramaniRajkumar32/ginLearn/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.Conn()
	// initializers.connect()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hi Veera",
		})
	})
	r.POST("/post", controllers.PostCreate)

	r.Run() // listen and serve ON env port
}
