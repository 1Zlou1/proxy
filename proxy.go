package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

const proxyAddr string = "localhost:9000"

var (
	counter            int    = 0
	firstInstanceHost  string = "localhost:8081"
	secondInstanceHost string = "localhost:8082"
)

func handleProxy(w http.ResponseWriter, r *http.Request) {

	textBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	text := string(textBytes)

	if counter == 0 {
		if _, err := http.Post(firstInstanceHost, "text/plain", bytes.NewBuffer([]byte(text))); err != nil {
			log.Fatalln(err)
		}
		counter++
		return
	}

	if counter == 1 {
		if _, err := http.Post(secondInstanceHost, "text/plain", bytes.NewBuffer([]byte(text))); err != nil {
			log.Fatalln(err)
		}
		counter--
		return
	}

}

func main() {
	http.HandleFunc("/", handleProxy)
	http.ListenAndServe(proxyAddr, nil)

}

