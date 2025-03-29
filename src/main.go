package main

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
)

type Telephone struct {
    Type        int    `json:"type"`
    Country     int    `json:"country"`
    Carrier     int    `json:"carrier"`
    Digits      int    `json:"digits"`
}

type Person struct {
    Identifier  int      `json:"identifier"`
    User        string   `json:"user"`
    FirstName   string   `json:"first_name"`
    LastName    string   `json:"last_name"`
    Name        string   `json:"name"`
    Phones      []Telephone `json:"phones"`
    Emails      []string `json:"emails"`
    DOB         string   `json:"dob"`
}

type Category struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    Info        string `json:"info"`
    Members     []int  `json:"members"`
}

func createEmptyPerson(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(Person{})
}

func createEmptyCategory(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(Category{})
}

func initializeServer() {
    router := mux.NewRouter()

    router.HandleFunc("/api/v1/person", createEmptyPerson).Methods("POST", "GET", "PUT", "DELETE")
    router.HandleFunc("/api/v1/category", createEmptyCategory).Methods("POST", "GET", "PUT", "DELETE")

    http.ListenAndServe(":6080", router)
}

func main() {
    initializeServer()
}
