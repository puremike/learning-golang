package main

import "log"

func main() {


    // loop through a set of numbers
    for i := 1; i <= 10; i++ {
        log.Println(i)
    }

    // loop through a slice of numbers
    log.Printf("......Loop through a slice of numbers......")
    m := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    for _, m := range m {
        log.Println(m)
    }


    // loop through a slice of string with their indexes
    log.Printf("........Loop through a slice of string with their indexes.......")
    s := []string {"apple", "banana", "cherry", "date", "elderberry"}

    for l, s := range s {
        log.Println(l, s)
    }

    // loop through a map of string
    log.Printf(".......Loop through a map of string..........")
    myMap := make(map[string]string)
    myMap["firstN"] = "Mike"
    myMap["lastN"] = "Johnson"

    for _, y := range myMap {
        log.Println(y)
    }

    // loop through a type struct
    log.Printf("........loop through a slice of type struct..........")

    type User struct {
        name, surname, email string
        age int
    }

    var users []User
    users = append(users, User{name: "John", surname: "Doe", email: "john.doe@example.com", age: 30})
    users = append(users, User{name: "Jane", surname: "Smith", email: "jane.smith@example.com", age: 28})
    users = append(users, User{name: "Bob", surname: "Johnson", email: "bob.johnson@example.com", age: 32})
    users = append(users, User{name: "Alice", surname: "Williams", email: "alice.williams@example.com", age: 25})
   
    for _, u := range users {
        log.Println(u.name, u.surname, u.email, u.age)
    }
	
}

