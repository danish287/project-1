package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/danish287/project-1/internal/dbClient"
	"github.com/danish287/project-1/internal/gethoroscope"
)

func main() {
	println("Server is running on port 8080")
	http.Handle("/", http.FileServer(http.Dir("web")))

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		var usrSunsign = r.FormValue("mySunsign")
		var reqDate = r.FormValue("date")
		reqHoroscope := gethoroscope.GetHoroscope(usrSunsign, reqDate)

		fmt.Fprint(w, "\n\n ", reqHoroscope)
	})

	http.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {

		var usrSunsign = r.FormValue("mySunsign")
		var usrName = r.FormValue("myName")
		var usrEmail = r.FormValue("myEmail")
		var usrPassword = r.FormValue("myPassword")
		answer := usrEmail + " " + usrName + " " + usrPassword + " " + usrSunsign
		fmt.Println(usrName)
		//reqHoroscope := gethoroscope.GetHoroscope(usrSunsign, reqDate)

		dbEmail := dbClient.FindEmail(usrEmail)
		if dbEmail == usrEmail {
			//will redirectback to index page for user to input fields again
			fmt.Println("Hi")
		} else {
			//dbClient.AddUser(usrName, usrEmail, usrPassword, usrSunsign)
			fmt.Println(r.URL.Path)
			fmt.Println("test")
			http.ServeFile(w, r, "./web/reject.html")
		}
		fmt.Println(answer)

	})

	http.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		fmt.Println("test")
		http.ServeFile(w, r, "./web/welcome.html")

	})

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		fmt.Println("test")
		http.ServeFile(w, r, "./web/index.html")

	})

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}

}
