package main

import (
	"fmt"
	"html/template"
	// "log"
	"net/http"
	// "os"
	// "time"
	"strings"
	// "syscall"
	// "os/signal"
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

func main() {

	mtmpl := template.Must(template.ParseFiles("templates/main.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/index.html")
	})

	http.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/signup.html")
		
	})

	http.HandleFunc("/newuser", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("we are here")
		fmt.Println(r.FormValue("mySunsign"))
		fmt.Println(r.FormValue("inputEmail"))

		if len(dbClient.FindEmail(r.FormValue("inputEmail"))) > 0 {
			http.Redirect(w,r, "/", http.StatusSeeOther)	
		} else{ 
			dbClient.AddUser(r.FormValue("inputName"), r.FormValue("inputEmail"), r.FormValue("inputPassword"), r.FormValue("mySunsign"), 0, false)
			w.Header().Set("Content-Type", "text/html")
			data := MainContent{
				Daily:   gethoroscope.GetHoroscope(r.FormValue("mySunsign"), "today"),
				Monthly: gethoroscope.GetHoroscope(r.FormValue("mySunsign"), "month"),
				Yearly:  gethoroscope.GetHoroscope(r.FormValue("mySunsign"), "year"),
				Zodaic:  r.FormValue("mySunsign"),
				UsrName: strings.Title(r.FormValue("inputName")),
			}
			err := mtmpl.Execute(w, data)
			if err != nil {
				fmt.Println(err)
		}

		}
	})

	http.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		usrEmail:= r.FormValue("inputEmail")
		usrPassword := r.FormValue("inputPassword")
		fmt.Println(usrEmail, usrPassword)
		fmt.Println(dbClient.Auth(usrEmail, usrPassword))
		fmt.Println("BLOCKED", !(dbClient.IsBlocked(usrEmail)))
		if !(dbClient.IsBlocked(usrEmail)) && dbClient.Auth(usrEmail, usrPassword){

			usrData := dbClient.FindUsr(usrEmail)
			w.Header().Set("Content-Type", "text/html")
			data := MainContent{
				Daily:   gethoroscope.GetHoroscope(usrData[2], "today"),
				Monthly: gethoroscope.GetHoroscope(usrData[2], "month"),
				Yearly:  gethoroscope.GetHoroscope(usrData[2], "year"),
				Zodaic:  usrData[2],
				UsrName: strings.Title(usrData[0]),
			}
			err := mtmpl.Execute(w, data)
			if err != nil {
					fmt.Println(err)
			}
			} else{
			dbClient.FailedAttempt(usrEmail) 
			http.Redirect(w,r, "/", http.StatusSeeOther)
		}
	})



	fmt.Println("Server is running on port 8081...")	
	err := http.ListenAndServe(":8081", nil)
	if err != nil{
		fmt.Println(err)
	}


	// terminate := make(chan error, 2)
	// interrupt := make(chan os.Signal, 1)
	// signal.Notify(interrupt, syscall.SIGINT)

	// go func() {
	// 	terminate <- http.ListenAndServe(":8080", nil)
	// }()

	// fmt.Println("Server is running on port 8081...")

	// for {
	// 	err := <-terminate
	// 	newSignal := <-interrupt
	// 	if (err != nil) || (newSignal != nil) {
	// 		fmt.Println("\nClosing application: ", newSignal)
	// 		os.Exit(0)

	// 	}

	// }



}
