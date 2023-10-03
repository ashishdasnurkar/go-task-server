package main

import (
	"encoding/base64"
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

	http.ListenAndServe("localhost:8080", router)
}

func syncTasks(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Reading body error: ", err)
	}
	fmt.Println(string(reqBody))
	taskArr, err := base64.StdEncoding.DecodeString(string(reqBody))
	if err != nil {
		log.Fatal("Reading body error: ", err)
	}

	var taskObj TaskType
	err = proto.Unmarshal(taskArr, &taskObj)
	if err != nil {
		log.Fatal("Unmarshaling error: ", err)
	}

	fmt.Println(taskObj)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
