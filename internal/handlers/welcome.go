package handlers

import "net/http"

// Welcome is a simple handler for the root endpoint.
func Welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Protein Bot from the Cloud!"))
}
