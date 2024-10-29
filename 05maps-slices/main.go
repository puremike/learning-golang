package main

import (
	"fmt"
	"log"
)

type User struct {
	firstName, lastName string
}

func main() {
	log.Println("MAPS............")
	myMap := make(map[string]string)
	mySecondMap := make(map[string]int)

	myMap["name"] = "Michael Egbinola"
	myMap["occupation"] = "Golang Developer"

	mySecondMap["age"] = 30
	mySecondMap["ninetynine"] = 99

	myThirdMap := make(map[string]User)
	
	user := User{
		firstName: "Foo",
		lastName: "Bar",
	}
	myThirdMap["firstN"] = user
	myThirdMap["lastN"] = user

	log.Println(myMap["name"], myMap["occupation"])
	log.Println(mySecondMap["age"], mySecondMap["ninetynine"])
	log.Println(myThirdMap["firstN"].firstName, myThirdMap["lastN"].lastName)

	log.Printf("\n")
	log.Println("SLICES...........")

	var colors []string
	colors = append(colors, "white", "black", "red", "blue")
	
	var age []int
	age = append(age, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	var newSlice = [2]string{"abc", "abcdef"}
		
	// Another way to declare arrays/slices
	colors2 := []string{"white", "black", "red", "blue"}
	age2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	newSlice2 := [10]int{5, 6, 7, 8, 9, 10, 3, 4, 2, 5}

	log.Println(colors)
	log.Println(colors2)
	log.Println(age)
	log.Println(newSlice)
	log.Println(age2)
	log.Println(newSlice2)
	log.Println(colors[0:3])
	log.Printf("\nFrom Checking Function")
	checking()

	obiOil, obiFood := sliceMethods(&colors, &age2)
	log.Println(fmt.Sprintf("The length of the slices are %d and %d respectively", obiOil, obiFood))
	
}


func checking() {
	var s []int
	primes := [6] int{2, 3, 4, 5, 6, 7}
	s = primes[3:5]
	log.Println("s before changing primes using pointer", s)
	p := &primes
	*p = [6] int{5, 6, 7, 8, 9, 10}
	s = primes[1:4]
	log.Println("s after changing primes using pointer", s)


}

func sliceMethods(c *[]string, d *[]int) (int, int){
	return len(*c), len(*d)
}


