package main

import (
	"net/http"
)

func main() {

	// Create a new request multiplexer
	// Take incoming requests and dispatch them to the matching handlers
	mux := http.NewServeMux()

	// Register the routes and handlers
	mux.Handle("/", &homeHandler{})
	//mux.Handle("/GNHF", &classHandler{})
	//mux.Handle("/gnhx", &classHandler{})

	// Run the server
	http.ListenAndServe(":42157", mux)
}

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
