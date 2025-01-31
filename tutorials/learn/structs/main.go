// type struct_name struct {
// 	member1 datatype;
// 	member2 datatype;
// 	member3 datatype;
// 	...
//   } 

package main

import (
	"fmt"
)

// declare a struct type Person with the following members: name, age, job and salary:
type Person struct {
	name string
	age int
	job string
	salary int
}

type gasEngine struct {
	mpg uint8
	gallons uint8
}

func main() {

	var myEngine gasEngine
	myEngine.mpg = 20
	fmt.Println(myEngine.mpg)

	var person1 Person
	var person2 Person

	// person1 specifications
	person1.name = "omer"
	person1.age = 22
	person1.job = "DevOps Engineer"
	person1.salary = 0

	// person2 specifications
	person2.name = "bobby"
	person2.age = 15
	person2.job = "Homeless"
	person2.salary = 50000

	fmt.Println("\nName: ", person1.name)
	fmt.Println("Age: ", person1.age)
	fmt.Println("Job: ", person1.job)
	fmt.Println("Salary: ", person1.salary)

	fmt.Println("\nName: ", person2.name)
	fmt.Println("Age: ", person2.age)
	fmt.Println("Job: ", person2.job)
	fmt.Println("Salary: ", person2.salary)

	printPerson(person1)
	printPerson(person2)
}

// to print it as a funtion
func printPerson(person Person) {
	fmt.Println("\nName: ", person.name)
	fmt.Println("Age: ", person.age)
	fmt.Println("Job: ", person.job)
	fmt.Println("Salary: ", person.salary)
  }