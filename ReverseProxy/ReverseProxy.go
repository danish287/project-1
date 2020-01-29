package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	myUrl, _ := url.Parse("http://localhost:8080")
	proxy := httputil.NewSingleHostReverseProxy(myUrl)
	http.ListenAndServe(":8081", proxy)

}
