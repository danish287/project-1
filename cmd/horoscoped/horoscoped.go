package main

import (
	"fmt"
	"net/http"

	"github.com/danish287/project-1/internal/gethoroscope"
)

func main() {
	println("Server is running on port 8080")
	http.Handle("/", http.FileServer(http.Dir("web")))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		var usrSunsign = r.FormValue("mySunsign")
		var reqDate = r.FormValue("date")
		reqHoroscope := gethoroscope.GetHoroscope(usrSunsign, reqDate)

		fmt.Fprint(w, "\n\n ", reqHoroscope)
	})

	http.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		var usrSunsign = r.FormValue("sunsign")
		var reqDate = r.FormValue("date")
		reqHoroscope := gethoroscope.GetHoroscope(usrSunsign, reqDate)
		fmt.Fprint(w, "\n\n ", reqHoroscope)
		fmt.Println("YOOOO")
	})

	http.ListenAndServe(":8080", nil)

}
