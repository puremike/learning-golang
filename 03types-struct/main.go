package main

import (
	"log"
	"time"
)

// IMPORTANT NOTE:
// Declaring variable or function with a starting Capital Letter makes the variable or functions publicly accessible. Example:
// func PubliclyAvailable () {} or var PubliclyAvailable - will be available publicly
// func privatelyAvailable () {} or var privatelyAvailable - won't be available publicly, but privately

// Instead of declaring many variables that are similar like this:

// var FirstName, LastName, Occupation, PhoneNumber string
// var BirthDate time.Time
// var Age int

// We can type struct to hold them together

type User struct {
FirstName, LastName, Occupation, PhoneNumber string
BirthDate time.Time
Age int }

func main() {

	user := User{
		FirstName: "Michael", 
		LastName: "Egbinola",
		// Occupation: "Software Engineer",
        // PhoneNumber: "+1 555-555-5555",
        // BirthDate: time.Date(1990, time.January, 15, 0, 0, 0, 0, time.UTC),
        // Age: 30,
	}

	log.Println(user.FirstName, user.LastName)
}





