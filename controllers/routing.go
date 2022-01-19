package controllers

import (
	"encoding/json"
	"net/http"
)

// RegisterController registers the controller
func RegisterController() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

// will encode Go objects into JSON 
func encodeResponseAsJSON(data interface{}, w http.ResponseWriter) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}