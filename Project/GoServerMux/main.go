package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	// Create a new Router 
	r := mux.NewRouter()
	// Method 1 
	// r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
	// 	vars := mux.Vars(r)
	// 	title := vars["title"]
	// 	page := vars["page"]
	// 	w.Header().Set("Content-Type", "application/json")
	// 	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	// }).Methods("GET")

	// Routes 
	r.HandleFunc("/books/{title}/page/{page}", getTitleAndPage()).Methods("GET")

	// Uses Port 8080 
	http.ListenAndServe(":8080", r)
}

// Method 2 
// Gets the title and page through http 
func getTitleAndPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	}

}
