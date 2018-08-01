package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	configpkg "github.com/rishikeshjadhav/GO/GoMongoDB/employeedb/webapi/config"
	daopkg "github.com/rishikeshjadhav/GO/GoMongoDB/employeedb/webapi/dao"
	modelspkg "github.com/rishikeshjadhav/GO/GoMongoDB/employeedb/webapi/models"
)

var config = configpkg.Config{}
var dao = daopkg.EmployeeDAO{}

// Function to create the employee
func Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API HIT -> Create")
	defer r.Body.Close()

	var emp modelspkg.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		respondWithError(w, http.StatusBadRequest, "Create - Invalid request payload")
		return
	}

	emp.ID = bson.NewObjectId()
	fmt.Printf("\nCreating user with Object Id as %s\n", emp.ID)

	if err := dao.Create(emp); err != nil {
		respondWithError(w, http.StatusBadRequest, "Create - Invalid request payload")
		return
	}
	fmt.Printf("\nCreated user with Object Id as %s and Code as %s\n", emp.ID, emp.Code)
	respondWithJson(w, http.StatusCreated, emp)
}

// Function to fetch employee(s)
func Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API HIT -> Get")
	emps, err := dao.Get()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Get - Invalid request payload")
		return
	}
	fmt.Printf("\nFound %d employees\n", len(emps))
	respondWithJson(w, http.StatusCreated, emps)
}

// Function to update the employee
func Update(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API HIT -> Update")
	defer r.Body.Close()

	var emp modelspkg.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		respondWithError(w, http.StatusBadRequest, "Update - Invalid request payload")
		return
	}

	fmt.Printf("\Updating user with Object Id as %s\n", emp.ID)
	if err := dao.Update(emp); err != nil {
		respondWithError(w, http.StatusBadRequest, "Update - Invalid request payload")
		return
	}
	fmt.Printf("\Updated user with Object Id as %s and Code as %s\n", emp.ID, emp.Code)
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// Function to delete the employee
func Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API HIT -> Delete")
	defer r.Body.Close()

	var emp modelspkg.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		respondWithError(w, http.StatusBadRequest, "Delete - Invalid request payload")
		return
	}

	fmt.Printf("\nDeleting user with Object Id as %s\n", emp.ID)
	if err := dao.Delete(emp); err != nil {
		respondWithError(w, http.StatusBadRequest, "Delete - Invalid request payload")
		return
	}
	fmt.Printf("\nDeleted user with Object Id as %s and Code as %s\n", emp.ID, emp.Code)
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// Function to Welcome the user
func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API HIT -> Welcome")
	fmt.Fprintln(w, "Welcome to Web API")
}

// Function to response with Json on error
func respondWithError(w http.ResponseWriter, code int, msg string) {
	fmt.Println("Error")
	fmt.Printf("\nMessage:\n %s\n", msg)
	respondWithJson(w, code, map[string]string{"error": msg})
}

// Function to reponse with Json on success
func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	fmt.Println("Success")
	response, _ := json.Marshal(payload)
	fmt.Printf("\nResponse:\n %s\n", response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Function init - Program initialize
func init() {
	fmt.Println("Initializing...")

	config.Read()

	fmt.Printf("\nFound server (%s) and database (%s)\n", config.Server, config.Database)

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()

	fmt.Println("Initialization Complete and Connection Established")
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

	fmt.Println("Listening on port 3000")

	// Listen and serve incoming requests
	log.Fatal(srv.ListenAndServe())

}
