package dao

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	mgo "gopkg.in/mgo.v2"
)

type EmployeeDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "employees"
)

// Function to establish a connection to MongoDB
func (m *EmployeeDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db := session.DB(m.Database)
}

// Function to create the employee
func Create(employee Employee) {
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
