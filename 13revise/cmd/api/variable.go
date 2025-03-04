package main

import (
	"fmt"
	"reflect"
)

func revisingVariable() (string, string) {
	// first way to declare a variable
	var name string = "Michael"

	//another way to declare a variable. Declare and pass the value one-liner
	var age, phoneNumber int = 26, 1234567890
	var isStudent, isPresent bool = true, false

	// another way to declare a variable
	lastName := "Egbinola"

	fmt.Printf("\nUsing Variables.........\n")

	fmt.Printf("My name is %s %s, I am a student: %t, and I am present: %t\n", name, lastName, isStudent, isPresent)
	fmt.Printf("I'm %d years old and my phone number is %d\n", age, phoneNumber)

	fmt.Println(reflect.TypeOf(lastName))

	return "Hello,", " " + name
}
