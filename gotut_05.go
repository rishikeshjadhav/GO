package main

import "fmt"

func main() {
	fmt.Println("Welcome to Go tutorial")

	// grades := make(map[string]float32)
	grades := map[string]float32

	grades["Rishikesh"] = 90
	grades["Nikhil"] = 80
	grades["Akshay"] = 60

	fmt.Println(grades)
}
