package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// Define handler functions for different routes
	indexHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is the index page")
	}

	dogHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is the dog page")
	}

	meHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is the me page.\nMy name is Eze Henry Chidera.")
	}

	// Register the handlers for specific routes using http.HandleFunc
	http.HandleFunc("/", indexHandler)   //index route
	http.HandleFunc("/dog/", dogHandler) //dog route
	http.HandleFunc("/me/", meHandler)   //me route

	//ListenAndServe on port ":8080" using the default ServeMux.
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
