package main

import (
	f "fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	f.Fprintln(w, "Apa kabar?")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		f.Fprintln(w, "halo!")
	})

	http.HandleFunc("/index", index)

	f.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}