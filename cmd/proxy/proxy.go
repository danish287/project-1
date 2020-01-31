package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

var (
	aPort  string
	bPort  string
	cPort  string
	myHost string
)

func main() {
	//getEnvVariables()
	runProxy()
}

func runProxy() {
	fmt.Println("Hi")
	//parse the url
	myURL, err := url.Parse("http://localhost:8080")
	//myHttps, _ := url.Parse("https://localhost")

	if err != nil {
		log.Fatal(err)
	}
	//initiate the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(myURL)
	//proxy2 := httputil.NewSingleHostReverseProxy(myHttps)
	//listen on given port
	//go http.ListenAndServeTLS(":443", "cert.pem", "key.pem", proxy)
	err = http.ListenAndServeTLS(":443", "cert.pem", "key.pem", proxy)
	//err = http.ListenAndServe(":80", proxy2)
	if err != nil {
		log.Fatal(err)
	}
}

func DirectProxy() {
	fmt.Println("Hi")
}

//getEnvVariables gets the environment variables for 3 ports and the host URL
func getEnvVariables() {
	aPort = os.Getenv("PORT_A")
	bPort = os.Getenv("PORT_B")
	cPort = os.Getenv("PORT_C")
	myHost = os.Getenv("MY_URL")
}
