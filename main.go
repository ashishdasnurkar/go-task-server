package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/sync", syncTasks).Methods("POST")

	fmt.Println("Server running on port 8080...")
	http.ListenAndServe("localhost:8080", router)
}

func syncTasks(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Reading body error: ", err)
	}
	fmt.Println(string(reqBody))

	var taskList TaskList
	err = proto.Unmarshal(reqBody, &taskList)
	if err != nil {
		log.Fatal("Unmarshaling error: ", err)
	}

	fmt.Printf("%+v", taskList)

	w.Write([]byte("Syncing Tasks"))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
