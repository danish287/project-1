package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type CLIHoroscopes struct {
	gorm.Model
	Name      string
	Horoscope []usrHoroscopes
}

var usrHoroscopes string

func AddHorosocpe(userName string, usrHoroscope string) {
	//returns DB object, specify the type of DB and the DB file
	db, err := gorm.Open("sqlite3", "myDB.db")
	var s []string
	s = append(s, usrHoroscope)
	fmt.Printf("%T", s)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	//Migrate the schema
	db.AutoMigrate(&CLIHoroscopes{})
	// db.Where()
	// Create
	db.Create(&CLIHoroscopes{Name: userName, Horoscope: {usrHoroscopes: usrHoroscope}})

	// Read
	var horos CLIHoroscopes

	db.First(&horos, 1)                   // find product with id 1
	db.First(&horos, "name = ?", "dania") // find product with code l1212

	// // Update - update product's price to 2000
	db.Model(&horos).Update("Name", "logan")

}

func main() {
	AddHorosocpe("dania", "test")
}
