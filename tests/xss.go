package main

import (
	"html/template"
	"log"
	"net/http"
)

type P struct {
	Data string
}


func set(w http.ResponseWriter, r *http.Request) {
	p := r.FormValue("data")
	log.Printf("received '%s'", p);
	t, _ := template.ParseFiles("xss.html")
	t.Execute(w, &P{Data:p})
}

func get(w http.ResponseWriter, r *http.Request) {
	p := "OMG awesome!"
	t, _ := template.ParseFiles("xss.html")
	t.Execute(w, &P{Data: p})
}

func main() {
	http.HandleFunc("/", get);
	http.HandleFunc("/save", set);
	http.ListenAndServe(":8989", nil)
}

