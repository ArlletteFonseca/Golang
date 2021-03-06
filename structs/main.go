package main

import "fmt"
type contactInfo struct {
	email string
	zipCode int
}
 type person struct {
	 //a struct of person with two fields, no commas no colons needed
	 firstName string
	 lastName string
	 //here you don't have to specify field name but you can also do contact contactInfo
	 contactInfo
 }
func main() {
	//here firstname is Alex and last name is Anderson, you are relying on order of fields
	/*alex := person{"Alex", "Anderson"}*/ 
	//below is another way of defining, and you can update using dot notation
	/* 	var alex person
	alex.firstName ="Alex"
	alex.lastName="Anderson"
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)
	//different way to print using %+v to print out "firstname" and "lastName" 
	fmt.Printf("%+v", alex) */
	//USE STRUCTS IF YOU HAVE A CERTAIN SET OF KEYS YOU KNOW YOU WILL BE WORKING WITH
	//values and keys do not have to be of same type
	jim := person {
		firstName: "Jim",
		lastName: "Party",
		contactInfo:contactInfo{
			email:"jim@gmail.com",
			zipCode: 94000,
		},
	}
	//& is an operator you write &and variable name give me the memory address of the value this variable is pointing at
	//shorcut here jim is a person struct and use *person in updateName
	jim.updateName("jimmy")
	jim.print()
	//test//
	}
	/* 
		turn address into a value with *address
		turn value into address with &value
		star in front of type pointer to a person
	 */
  func (pointerToPerson *person) updateName(newFirstName string) {
	  //* operator give me the value this memory address is pointing at
	  (*pointerToPerson ).firstName = newFirstName
  }
	//create function that takes receiver
	func (p person) print() {
		fmt.Printf("%+v", p)
	}