package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	runProxy()
}

func runProxy() {
	//parse the url
	myURL, err := url.Parse("http://localhost:8080")

	if err != nil {
		log.Fatal(err)
	}
	//initiate the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(myURL)

	fmt.Println("Server is running on port 443")

	//listen on given ports
	err = http.ListenAndServeTLS(":443", "cert.pem", "key.pem", proxy)
	if err != nil {
		log.Fatal(err)
	}

}
