package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

const port = ":8081"

func main() {
	log.Printf("Starting server on %s for testing HTTP...\n", port)

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/test/{data}", GetHandler)
	router.Post("/test", PostHandler)
	router.Patch("/test", PatchHandler)
	router.Delete("/test/{data}", DeleteHandler)
	router.Put("/test", PutHandler)

	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal(err)
	}
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handler Get request on %s port", port)

	data := chi.URLParam(r, "data")

	result := map[string]string{
		"data":            data,
		"service_on_port": port,
	}

	response, err := json.Marshal(result)
	if err != nil {
		buildResponse(w, http.StatusInternalServerError, nil)
		return
	}
	buildResponse(w, http.StatusOK, response)
}

type PostData struct {
	Data string `json:"data"`
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handler Post request on %s port", port)

	postdata := &PostData{}
	err := json.NewDecoder(r.Body).Decode(postdata)
	if err != nil {
		buildResponse(w, http.StatusInternalServerError, nil)
		return
	}

	result := map[string]string{
		"data":            postdata.Data,
		"service_on_port": port,
	}

	response, err := json.Marshal(result)
	if err != nil {
		buildResponse(w, http.StatusInternalServerError, nil)
		return
	}
	buildResponse(w, http.StatusCreated, response)
}

type PatchData struct {
	Data string `json:"data"`
}

func PatchHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handler Patch request on %s port", port)

	patchdata := &PatchData{}
	err := json.NewDecoder(r.Body).Decode(patchdata)
	if err != nil {
		buildResponse(w, http.StatusInternalServerError, nil)
		return
	}

	result := map[string]string{
		"data":            patchdata.Data,
		"service_on_port": port,
	}

	response, err := json.Marshal(result)
	if err != nil {
		buildResponse(w, http.StatusInternalServerError, nil)
		return
	}
	buildResponse(w, http.StatusOK, response)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handler Delete request on %s port", port)

	deletedata := chi.URLParam(r, "data")

	result := map[string]string{
		"data":            deletedata,
		"service_on_port": port,
	}

	response, err := json.Marshal(result)
	if err != nil {
		buildResponse(w, http.StatusInternalServerError, nil)
		return
	}
	buildResponse(w, http.StatusOK, response)
}

type PutData struct {
	Data string `json:"data"`
}

func PutHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handler Put request on %s port", port)

	putdata := &PutData{}
	err := json.NewDecoder(r.Body).Decode(putdata)
	if err != nil {
		buildResponse(w, http.StatusInternalServerError, nil)
		return
	}

	result := map[string]string{
		"data":            putdata.Data,
		"service_on_port": port,
	}

	response, err := json.Marshal(result)
	if err != nil {
		buildResponse(w, http.StatusInternalServerError, nil)
		return
	}
	buildResponse(w, http.StatusOK, response)
}

func buildResponse(w http.ResponseWriter, statusCode int, body []byte) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	w.Write(body)
}
