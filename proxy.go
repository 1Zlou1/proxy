package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	target := "http://localhost:8081"
	proxy, err := createReverseProxy(target)
	if err != nil {
		log.Fatalln(err)
	}

	http.Handle("/", proxy)
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func createReverseProxy(target string) (http.Handler, error) {
	targetURL, err := url.Parse(target)
	if err != nil {
		return nil, err
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	return proxy, nil
}
