package main

import "fmt"

type UserDetails struct {
	name, occupation, email string
	age, number             int
}

func useTypeStruct() {
	user := UserDetails{
		name:       "John Doe",
		occupation: "Software Developer",
		age:        30,
		number:     1234567890,
		email:      "john@example.com",
	}

	fmt.Printf("\nUsing Type Struct...........\n")

	fmt.Printf("My name is %s and I am a %s\n", user.name, user.occupation)
	fmt.Printf("I'm %d years old and my phone number is %d\n", user.age, user.number)
	fmt.Printf("My email address is %s\n", user.email)
}
