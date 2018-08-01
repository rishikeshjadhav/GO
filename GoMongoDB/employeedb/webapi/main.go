package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Function to create the employee
func Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var emp Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		respond
	}

	fmt.Fprintln(w, "Create: not implemented yet")
}

// Function to fetch employee(s)
func Get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Get: not implemented yet")
}

// Function to update the employee
func Update(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Update: not implemented yet")
}

// Function to delete the employee
func Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Delete: not implemented yet")
}

// Function to Welcome the user
func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome: not implemented yet")
}

// Function main - Program start
func main() {
	fmt.Println("Welcome to Web API")

	// Create new handler
	r := mux.NewRouter()

	// Define handler attributes
	// r.Host("localhost")
	// r.PathPrefix("/api/")
	// r.Methods("GET", "POST", "PUT", "DELETE")
	// r.Schemes("https")
	// r.Headers("X-Request-With")

	// Define routes for handler
	r.HandleFunc("/", Welcome).Methods("GET")
	r.HandleFunc("/Create", Create).Methods("POST")
	r.HandleFunc("/Get", Get).Methods("GET")
	r.HandleFunc("/Update", Update).Methods("PUT")
	r.HandleFunc("/Delete", Delete).Methods("DELETE")

	// http.ListenAndServe(":3000", r)

	// Create server with handler and timeouts
	srv := &http.Server{
		Handler:      r,
		Addr:         "localhost:3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Listen and serve incoming requests
	log.Fatal(srv.ListenAndServe())

}
