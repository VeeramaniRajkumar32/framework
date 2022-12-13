package initializers

import (
	"fmt"

	"github.com/VeeramaniRajkumar32/ginLearn/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

const (
	USER     = "root"
	PASSWORD = ""
	DATABASE = "blog"
)

func Conn() {
	dsn := fmt.Sprintf(
		"%s:%s@(localhost)/%s?charset=utf8&parseTime=True&loc=Local",
		USER, PASSWORD, DATABASE,
	)

	DB, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Database connected successfully")
	// defer DB.Close()
	DB.AutoMigrate(&models.Post{})

}
