package main

import (
	"fmt"
)

type User struct {
	email string
	username string
	age int32
}

func (u User) Email() string {
	return u.email
}

func (u *User) updateEmail(email string) {
	u.email = email
}

func main() {
	user := User {
		email: "omer@gmail.com",
		username: "omer",
		age: 22,
	}
	user.updateEmail("blah@gmail.com")
	fmt.Println(user.Email())
}



// type Person struct {
// 	firstName string
// 	lastName string
// 	Age int32
// }

// func (p Person) fullName() string {
// 	return p.firstName + " " + p.lastName
// }

// func main() {
// 	person1 := Person {
// 		firstName: "Omer",
// 		lastName: "Revach",
// 		Age: 22,
// 	}

// 	fmt.Println("First Name: ", person1.firstName)
// 	fmt.Println("Last Name: ", person1.lastName)
// 	fmt.Println("Full Name: ", person1.fullName())
// 	fmt.Println("Age: ", person1.Age)
// }

