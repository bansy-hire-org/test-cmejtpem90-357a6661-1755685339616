package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type License struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var licenses = []License{
	{ID: 1, Name: "License A"},
	{ID: 2, Name: "License B"},
	{ID: 3, Name: "License C"},
}

func getLicenses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(licenses)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/licenses", getLicenses).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
