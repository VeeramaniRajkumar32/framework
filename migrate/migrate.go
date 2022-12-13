package main

import (
	"github.com/VeeramaniRajkumar32/ginLearn/initializers"
	"github.com/VeeramaniRajkumar32/ginLearn/models"
)

func init() {
	initializers.Conn()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})

	// type user struct {
	// 	id        int
	// 	username  string
	// 	password  string
	// 	createdAt time.Time
	// }

	// rows, err := initializers.DB.AutoMigrate(&user)
	// // Query(`SELECT * FROM users`)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()

	// var users []user
	// for rows.Next() {
	// 	var u user

	// 	err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	users = append(users, u)
	// }
	// if err := rows.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%#v", users)
	// }
}
