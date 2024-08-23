package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type student struct {
	Name  string `json:"name"`
	Roll  int    `json:"roll"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func StuDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	st := student{}
	fmt.Println(st, "Before calling a API")
	json.NewDecoder(r.Body).Decode(&st)
	fmt.Println(st, "After calling a API")
}

func Greeting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World"))
}

func SumOfDigit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	type arithmetic struct {
		A int `json:"a"`
		B int `json:"b"`
	}
	at := arithmetic{}
	json.NewDecoder(r.Body).Decode(&at)
	C := at.A + at.B
	s := "Sum Of " + strconv.Itoa(at.A) + " And " + strconv.Itoa(at.B) + " = " + strconv.Itoa(C)
	json.NewEncoder(w).Encode(&s)
}
