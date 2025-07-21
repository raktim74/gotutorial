package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// handleReq()
	router := mux.NewRouter()
	router.HandleFunc("/item-data", getItems).Methods("GET")
	http.ListenAndServe(":9000", router)
}

func handleReq() {
	// http.HandleFunc("/", homepage)
	log.Fatal(http.ListenAndServe(":9000", nil))
	http.ListenAndServe(":9000", nil)
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Homepage")
}

type Item struct {
	ID         string       `json:"id"`
	Name       string       `json:"name"`
	Tags       []string     `json:tags`
	Attributes []Attributes `json:characteristics`
}

type Attributes struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

var items = []Item{
	{ID: "1", Name: "Raktim", Tags: []string{"go", "python", "nodejs"}, Attributes: []Attributes{
		{Name: "Complexion", Value: "Fair"},
		{Name: "T-Shirt", Value: "M"},
	}},
	{ID: "2", Name: "Baba", Tags: []string{"sql", "etl", "mongo"}, Attributes: []Attributes{
		{Name: "Complexion", Value: "Dark"},
		{Name: "T-Shirt", Value: "XL"},
	}},
}

func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}
