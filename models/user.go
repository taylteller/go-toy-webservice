package models

import (
	"errors"
	"fmt"
)

// procedural go

// User defines a new user
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User // a slice that holds pointers to User objects
	nextID = 1     // no data type bc this initialization is at package level; it's implicit with the 1
)

// GetUsers returns the slice of User pointers
func GetUsers() []*User {
	return users
}

// AddUser creates a new User
func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("New User must not include id, or it must be set to 0")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

// GetUserByID fetches a user by id
func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("No user found with ID '%v'", id)
}

// UpdateUser updates the user that matches the user id with the rest of the content in the passed in User object
func UpdateUser(u User) (User, error) {
	for i, val := range users {
		if val.ID == u.ID {
			users[i] = &u // we want the address of the passed in u User, since users slice contains pointers
			return u, nil
		}
	}
	return User{}, fmt.Errorf("No user found with ID '%v'", u.ID)
}

// RemoveUserByID takes an id and removes the user associated with it
func RemoveUserByID(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("No user found with ID '%v'", id)
}