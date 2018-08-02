package dao

import (
	"fmt"
	"log"

	"github.com/rishikeshjadhav/GO/GoMongoDB/employeedb/webapi/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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
	fmt.Printf("\nStarting connection with server %s\n", m.Server)
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nConnecting to database %s\n", m.Database)
	db := session.DB(m.Database)
	if db == nil {
		fmt.Println("DB is nil")
	}
	fmt.Println(db)
}

// Function to create the employee
func (m *EmployeeDAO) Create(employee models.Employee) error {
	fmt.Println("DAO HIT -> Create")
	if db == nil {
		fmt.Println("DB is nil")

		m.Connect()

		fmt.Println("Created connection with DB")
	}
	err := db.C(COLLECTION).Insert(&employee)
	return err
}

// Function to fetch employee(s)
func (m *EmployeeDAO) Get() ([]models.Employee, error) {
	fmt.Println("DAO HIT -> Get")
	var employees []models.Employee
	err := db.C(COLLECTION).Find(bson.M{}).All(&employees)
	return employees, err
}

// Function to update the employee
func (m *EmployeeDAO) Update(employee models.Employee) error {
	fmt.Println("DAO HIT -> Update")
	err := db.C(COLLECTION).UpdateId(employee.ID, &employee)
	return err
}

// Function to delete the employee
func (m *EmployeeDAO) Delete(employee models.Employee) error {
	fmt.Println("DAO HIT -> Delete")
	err := db.C(COLLECTION).Remove(&employee)
	return err
}
