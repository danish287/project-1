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

	//myHttps, _ := url.Parse("https://localhost")
	//proxy2 := httputil.NewSingleHostReverseProxy(myHttps)
	//go http.ListenAndServeTLS(":443", "cert.pem", "key.pem", proxy)
	//err = http.ListenAndServe(":80", proxy2)
}

//getPort checks if there are any avaiable ports to run our server on from the three ports we want to distrubute to, else logs error that there are no avilable servers
// func getPort() string {

// 	conn, err := net.Listen("tcp", aPort)
// 	if err == nil {
// 		conn.Close()
// 		return aPort
// 	}
// 	conn, err = net.Listen("tcp", bPort)
// 	if err == nil {
// 		conn.Close()
// 		return bPort
// 	}

// 	conn, err = net.Listen("tcp", cPort)
// 	if err == nil {
// 		conn.Close()
// 		return cPort
// 	}

// 	return "No avialble ports"
// }
