package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name  string
	Phone string
}

func main() {
	fmt.Println("Welcome to MongoDB connection with GO Lang")

	session, err := mgo.Dial("localhost:27017")

	if err != nil {
		panic(err)
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("HelloGoWorldDB").C("people")

	c.Insert(&Person{"Rishikesh", "1234567890"},
		&Person{"Nikhil", "0987654321"})

	if err != nil {
		log.Fatal(err)
	}

	result := []Person{}

	c.Find(bson.M{"name": "Nikhil"}).All(&result)

	if err != nil {
		log.Fatal(err)
	}

	for _, item := range result {
		fmt.Printf("\n%s (%s): ", item.Name, item.Phone)
	}

	fmt.Println("Completed")
}
