package main

import (
	"net/http"
)

func main() {
	// Create a new request multiplexer
	mux := http.NewServeMux()

	// Register the specific handlers
	mux.Handle("/home", &homeHandler{})
	mux.Handle("/part-number-decoder", &pnHandler{})
	mux.Handle("/iso-decoder", &isoHandler{})

	// Set the default handler for unmatched paths
	mux.Handle("/", &productNumberHandler{})

	// Run the server
	http.ListenAndServe(":42157", mux)
}

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

type productNumberHandler struct{}

func (h *productNumberHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/part-number-decoder" || r.URL.Path == "/iso-decoder" || r.URL.Path == "/home" {

	} else {
		http.ServeFile(w, r, "product-number-decoder.html")
	}

}

type pnHandler struct{}

func (h *pnHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "part-number-decoder.html")
}

type isoHandler struct{}

func (h isoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "iso-decoder.html")
}
