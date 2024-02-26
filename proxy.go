package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var (
	firstaddr  = "http://localhost:8081"
	secondaddr = "http://localhost:8082"
	count      = 0
)

func main() {

	http.Handle("/", Randserv())
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

func Randserv() http.Handler {
	if count == 0 {
		proxy, err := createReverseProxy(firstaddr)
		if err != nil {
			log.Fatalln(err)
		}
		count++
		return proxy
	}

	if count == 1 {
		proxy, err := createReverseProxy(secondaddr)
		if err != nil {
			log.Fatalln(err)
		}
		count--
		return proxy
	}
	return nil
}
