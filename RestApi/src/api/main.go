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
	r.HandleFunc("/Student", router.StuDetails).Methods("POST")
	r.HandleFunc("/Arithmetic", router.SumOfDigit).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}
