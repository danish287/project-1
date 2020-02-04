package dbClient

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//UsrLogin stores the name, email, and password for user authentication
type UsrLogin struct {
	gorm.Model
	Name       string
	Email      string
	Password   string
	usrSunsign string
}

//AddUser adds a user (name, email, and password) to database for user authertication
func AddUser(userName string, usrEmail string, usrPassword string, usrSunsign string) {
	//returns DB object, specify the type of DB and the DB file
	db, err := gorm.Open("sqlite3", "myDB.db")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	//Migrate the schema
	db.AutoMigrate(&UsrLogin{})
	db.Create(&UsrLogin{Name: userName, Email: usrEmail, Password: usrPassword})
}

//FindEmail checks if given email is on out database
func FindEmail(userEmail string) string {
	db, err := gorm.Open("sqlite3", "myDB.db")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	var currUsr UsrLogin
	db.Find(&currUsr, "email = ?", userEmail)
	return currUsr.Email
}
