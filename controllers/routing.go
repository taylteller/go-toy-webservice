package controllers

import "net/http"

// RegisterController registers the controller
func RegisterController() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
