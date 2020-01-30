package main

import (
	"fmt"
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
	getEnvVariables()
	runProxy()
}

func runProxy() {
	//parse the url
	myURL, _ := url.Parse(myHost)
	//initiate the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(myURL)
	//listen on given port
	http.ListenAndServe(aPort, proxy)
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
