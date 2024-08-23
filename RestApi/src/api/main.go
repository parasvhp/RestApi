package main

import (
	"github.com/gorilla/mux"
	"github.com/paras/api/router"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/greeting", router.Greeting).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
