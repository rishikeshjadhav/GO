package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type NewsAppPage struct {
	Title string
	News  string
}

func aggHandler(w http.ResponseWriter, r *http.Request) {
	p := NewsAppPage{Title: "News App Title", News: "This is a news"}
	t, _ := template.ParseFiles("agghtml.html")
	t.Execute(w, p)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hey, HTML from go tut</h1>")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg", aggHandler)
	http.ListenAndServe(":8000", nil)
}
