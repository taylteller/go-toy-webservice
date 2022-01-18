package models

// procedural go

// User defines a new user
type User struct {
	ID int
	FirstName string
	LastName string
}

var (
	users []*User // a slice that holds pointers to User objects
	nextID = 1 // no data type bc this initialization is at package level; it's implicit with the 1
)

// GetUsers returns the slice of User pointers
func GetUsers() []*User {
	return users
}

// AddUser creates a new User
func AddUser(u User) (User, error) {
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}