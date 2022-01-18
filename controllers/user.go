package controllers

import (
	"net/http"
	"regexp"
)

// object-oriented go; userController object has methods associated with it

type userController struct {
	userIDPattern *regexp.Regexp
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("in user controller"))
}

// newUserController is a constructor fn for userController
func newUserController() *userController {
	// return a pointer instead of the object itself so we can avoid a pointless copy operation
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`), // paths that are /users/[a number]/
	}
}
