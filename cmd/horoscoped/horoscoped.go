package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/danish287/project-1/internal/dbClient"
	"github.com/danish287/project-1/internal/gethoroscope"
)

type MainContent struct {
	Daily   string
	Monthly string
	Yearly  string
	Zodaic  string
	UsrName string
}

type Welcome struct {
	MyName string
}

func main() {
	println("Server is running on port 8081")

	// logging anything that happens in this file
	currTime := time.Now()
	path := "logs/horoscoped/" + currTime.Format("01-02-2006") + ".log"
	file, _ := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	log.SetOutput(file)

	mtmpl := template.Must(template.ParseFiles("web/main.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/index.html")
	})

	http.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/signup.html")
	})

	http.HandleFunc("/locked", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/locked.html")
	})

	http.HandleFunc("/main", func(w http.ResponseWriter, r *http.Request) {
		var usrEmail = r.FormValue("inputEmail")
		var usrPassword = r.FormValue("inputPassword")
		fmt.Println(usrPassword)
		if usrPassword == "" {
			fmt.Println("yes")
		}
		fmt.Println(usrEmail)
		dbEmail := dbClient.FindEmail(usrEmail)
		isUsr := dbClient.Auth(usrEmail, usrPassword)
		isBlocked := !(dbClient.IsBlocked(usrEmail))
		fmt.Println(dbEmail)
		if (dbEmail == usrEmail) && isUsr && isBlocked {
			w.Header().Set("Content-Type", "text/html")
			data := MainContent{
				Daily:   gethoroscope.GetHoroscope("aquarius", "today"),
				Monthly: gethoroscope.GetHoroscope("aquarius", "month"),
				Yearly:  gethoroscope.GetHoroscope("aquarius", "year"),
				Zodaic:  "leo",
				UsrName: "Dania",
			}
			mtmpl.Execute(w, data)
		} else if !(isBlocked) {
			http.ServeFile(w, r, "./web/locked.html")
		} else {
			http.ServeFile(w, r, "./web/index.html")
		}
	})

	err := http.ListenAndServe(":8081", nil)

	if err != nil {
		log.Fatal(err)
	}

}
