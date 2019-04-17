package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Greeting struct {
	ID       string `json:"id,omitempty"`
	Language string `json:"language,omitempty"`
	Message  string `json:"message,omitempty"`
}

var greetings []Greeting

func GetGreetings(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(greetings)
}

func GetGreeting(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range greetings {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Greeting{})
}

func CreateGreeting(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var greeting Greeting
	_ = json.NewDecoder(r.Body).Decode(&greeting)
	greeting.ID = params["id"]
	greetings = append(greetings, greeting)
	json.NewEncoder(w).Encode(greeting)
}
func DeleteGreeting(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range greetings {
		if item.ID == params["id"] {
			greetings = append(greetings[:index], greetings[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(greetings)
	}
}

func main() {
	router := mux.NewRouter()

	greetings = append(greetings, Greeting{ID: "1", Language: "en", Message: "Hello"})
	greetings = append(greetings, Greeting{ID: "2", Language: "es", Message: "Hola"})

	router.HandleFunc("/greeting", GetGreetings).Methods("GET")
	router.HandleFunc("/greeting/{id}", GetGreeting)
	router.HandleFunc("/greeting/{id}", CreateGreeting).Methods("POST")
	router.HandleFunc("/greeting/{id}", DeleteGreeting).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
