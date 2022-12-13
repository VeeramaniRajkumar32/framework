package initializers

// import (
// 	"database/sql"
// 	"log"
// )

// var DB *sql.DB

// const (
// 	USER     = "root"
// 	PASSWORD = ""
// 	DATABASE = "blog"
// )

// func connect() {
// 	var err error
// 	dsn := USER + ":" + PASSWORD + "@(localhost)/" + DATABASE + "?charset=utf8&parseTime=True&loc=Local"

// 	DB, err := sql.Open("mysql", (dsn))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err := DB.Ping(); err != nil {
// 		log.Fatal(err)
// 	}
// 	defer DB.Close()

// 	// {

// 	// 	query := `SELECT * FROM course`

// 	// 	// Executes the SQL query in our database. Check err to ensure there was no error.
// 	// 	// _, err := DB.Exec(query)
// 	// 	if _, err := DB.Exec(query); err != nil {
// 	// 		log.Fatal(err)
// 	// 	}
// 	// }
// }
