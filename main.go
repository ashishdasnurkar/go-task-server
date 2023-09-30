package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/sync", syncTasks).Methods("POST")

	http.ListenAndServe("localhost:8080", router)
}

func syncTasks(w http.ResponseWriter, r *http.Request) {
	
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
