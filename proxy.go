package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var (
	firstAddr  = "http://localhost:8081"
	secondAddr = "http://localhost:8082"
	count      = 0
)

func main() {
	http.HandleFunc("/", Randserv)
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func createReverseProxy(target string) (http.Handler, error) {
	targetURL, err := url.Parse(target)
	if err != nil {
		return nil, err
	}
	return httputil.NewSingleHostReverseProxy(targetURL), nil
}

func Randserv(res http.ResponseWriter, req *http.Request) {
	if count == 0 {
		proxy, err := createReverseProxy(firstAddr)
		if err != nil {
			log.Println(err)
			return
		}
		count++
		proxy.ServeHTTP(res, req)
	}

	if count == 1 {
		proxy, err := createReverseProxy(secondAddr)
		if err != nil {
			log.Println(err)
			return
		}
		count--
		proxy.ServeHTTP(res, req)
	}
}
