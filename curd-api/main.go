package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:id, omitempty`
	FirstName string   `json:firstname, omitempty`
	LastName  string   `json:lastname, omitempty`
	Address   *Address `json:address, omitempty`
}

type Address struct {
	City  string `json:city, omitempty`
	State string `json:state, omitempty`
}

var people []Person

func GetPeopleEndPoint(w http.ResponseWriter, req *http.Request) {
	fmt.Println("EndPoint Hit: All People")
	json.NewEncoder(w).Encode(people)
}

func GetPersonEndPoint(w http.ResponseWriter, req *http.Request) {
	fmt.Println("EndPoint Hit: Get One Person")
	params := mux.Vars(req)
	for _, person := range people {
		if person.ID == params["id"] {
			json.NewEncoder(w).Encode(person)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func CreatePersonEndPoint(w http.ResponseWriter, req *http.Request) {
	fmt.Println("EndPoint Hit: Create One Person")
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func DeletePersonEndPoint(w http.ResponseWriter, req *http.Request) {
	fmt.Println("EndPoint Hit: Delete One Person")
	params := mux.Vars(req)
	for index, person := range people {
		if person.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
		}
	}
	json.NewEncoder(w).Encode(people)
}

func handleRequest() {
	router := mux.NewRouter().StrictSlash(true)

	people = append(people, Person{ID: "1", FirstName: "Azmul", LastName: "Hossain", Address: &Address{City: "Dhaka", State: "Bangladesh"}})
	people = append(people, Person{ID: "2", FirstName: "Nazmul", LastName: "Ahmed"})

	// endpoints
	router.HandleFunc("/peoples", GetPeopleEndPoint).Methods("GET")
	router.HandleFunc("/peoples/{id}", GetPersonEndPoint).Methods("GET")
	router.HandleFunc("/peoples/{id}", CreatePersonEndPoint).Methods("POST")
	router.HandleFunc("/peoples/{id}", DeletePersonEndPoint).Methods("DELETE")

	// server create
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	handleRequest()
}
