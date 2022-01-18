package main

import (
	"fmt"
	"net/http"

	"github.com/taylteller/go-webservice/controllers"
)

func main() {
	// pointersCheatSheet()

	controllers.RegisterController()
	http.ListenAndServe(":3000", nil) // front controller/back controller pattern
}

func pointersCheatSheet() {
	// not a useful function except as a reference on how go pointers work, with a house analogy
	var firstName *string = new(string)
	*firstName = "test"                    // dereference the pointer - enter the house to put something in
	fmt.Println("Hello hello", *firstName) // dereference the pointer - enter the house to get the contents

	// addres of operator: &
	lastName := "Foo"
	ptr := &lastName //we're saying I want ptr to be a pointer. So I want to assign to ptr the ADDRESS of lastName

	fmt.Println(ptr, *ptr) // will print [some memory address] Foo
	// so here I'm getting the address first. Second I'm actually entering the house to retrieve the contents.

	lastName = "Boo"                 // assign this variable, which will stay at the same address, a new value
	fmt.Println(ptr, *ptr)           // will print [the same memory address] Boo
	fmt.Println(ptr, lastName)       // will print the same thing as the line above
	fmt.Println(&lastName, lastName) // will print the same thing as the line above

	middleName := lastName // becomes Boo

	fmt.Println("middleName", middleName)

	lastName = "new last name"
	fmt.Println("middleName", middleName) // will still be Boo
	fmt.Println("lastName", lastName)     // is now new last name
	// so basically a variable can hold EITHER the contents of the house (lastName) OR the address of the house (ptr)
	// if your variable contains the contents, use &variable to get the address
	// if your variable contains the address, use *variable to get the contents
}
