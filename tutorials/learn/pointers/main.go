package main

import (
	"fmt"
)

// When to use pointers
// 1. When we need to update the state

type User struct {
	email string
	username string
	age int32
}

func (u User) Email() string {
	return u.email
}

func (u *User) UpdateEmail(email string) { // if i didnt add the * in User then the email wouldnt get updated
	u.email = email
}

func changeValue(str *string) {
	*str = "blah"
} 

func changedValue2(str string) {
	str = "change"
}

func main() {
	user := User {
		email: "omer@gmail.com",
		username: "omer",
		age: 22,
	}
	user.UpdateEmail("blah@gmail.com")
	fmt.Println(user.Email())

	x := 7
	
	y := &x // now y points to the memry location of x
	
	*y = 5 // if i point to it i can change x value

	fmt.Println(x) // The & tells me where the value of x is stored

	toChange := "Hello"

	changedValue2(toChange) // wont change 
	fmt.Println(toChange)

	changeValue(&toChange) // This will change it because its pointing and changing its location in memory 
	fmt.Println(toChange)
}

