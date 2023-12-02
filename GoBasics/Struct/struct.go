package main

import (
	"fmt"
	"strconv"
)

// Go’s structs are typed collections of fields.
// They’re useful for grouping data together to form records.
// Declare Struct
type User struct {
	userName string
	age uint8
	email string
	password string
}

// createUser - creates a User 
func createUser(name string) *User {
	user := User{userName: name}
	user.age = 30
	fmt.Println("User address :",&user)
	return &user
}

// Method for User - Pass by value
// func (user User) introduce() string {
// 	fmt.Println("Address of users name :",&user.userName)
// 	user.email = "joff@gmail.com"
// 	return "\nHey my name is " + user.userName + " and my age is " + strconv.FormatUint(uint64(user.age), 10)
// } 

// Method for User - Pass by reference
func (user *User) introduce() string {
	fmt.Println("Address of users name :",&user.userName)
	user.email = "joff@gmail.com"
	return "\nHey my name is " + user.userName + " and my age is " + strconv.FormatUint(uint64(user.age), 10)
} 

func main(){
	// Add data into the User 
	user := User{ age: 21 }
	fmt.Println("User :",user)
	user.userName = "Joyfree Dias"
	user.email = "joyfree@gmail.com"
	user.password = "Joke"
	fmt.Println("User :",user)

	// create a User 
	userPointer := createUser("Smith Jone")
	userPointer.email = "smithjone@gmail.com"
	fmt.Println("User Pointer :",userPointer)

	// Another method to declare struct 
	dog := struct {
		name string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println("Dog name :",dog.name)
	fmt.Println("Dog :",dog)

	// Call User method 'introduce' 
	fmt.Println(user.introduce())
	fmt.Println("Address of users name :",&user.userName)
	fmt.Println("User email :",user.email)
}