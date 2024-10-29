// package main

// import "fmt"

// // Animal defines the interface for type Animal
// type Animal interface {
// 	Says() string
// 	NumberOfLegs() int
// }

// // Dog defines the dog type
// type Dog struct {
// 	Name  string
// 	Breed string
// }

// // Gorilla defines the Gorilla type
// type Gorilla struct {
// 	Name          string
// 	Color         string
// 	NumberOfTeeth int
// }

// func main() {
// 	dog := Dog{
// 		"Samson", "German Shepherd",
// 	}

// 	// We can pass dog to PrintInfo(), since the Dog type implements the Animal interface by having all of the
// 	// necessary functions. The parameter is passed as a pointer since the functions for the type have pointer
// 	// receivers (which is the norm. See https://tour.golang.org/methods/4 and
// 	// https://tour.golang.org/methods/8 for more details).
// 	PrintInfo(&dog)

// 	gorilla := Gorilla{
// 		"Jock", "grey", 38,
// 	}

// 	// We can also pass gorilla to PrintInfo(), since the Gorilla type implements the Animal interface by having all of the
// 	// necessary functions. The parameter is passed as a pointer since the functions for the type have pointer
// 	// receivers (which is the norm. See https://tour.golang.org/methods/4 and
// 	// https://tour.golang.org/methods/8 for more details).
// 	PrintInfo(&gorilla)
// }

// func PrintInfo(a Animal) {
// 	fmt.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
// }

// // Says has a receiver of type *Dog, so it satisfies part of the interface requirements for Animal
// // for the Dog type
// func (d *Dog) Says() string {
// 	return "Woof"
// }

// // NumberOfLegs satisfies the rest of the Animal interface requirements for the Dog type
// func (d *Dog) NumberOfLegs() int {
// 	return 4
// }

// // Says has a receiver of type *Gorilla, so it satisfies part of the interface requirements for Animal
// // for the Gorilla type
// func (d *Gorilla) Says() string {
// 	return "Ugh"
// }

// // NumberOfLegs satisfies the rest of the Animal interface requirements for the Gorilla type
// func (d *Gorilla) NumberOfLegs() int {
// 	return 2
// }

package main

import "fmt"

type Animal interface {
	says() string
	numberOfLegs() int
}

type Dog struct {
    name, breed string
}

type Gorilla struct {
	name, color string
	numberOfTeeth int
}

func main() {


	dog := Dog{
        name:  "Samson",
        breed: "German Shepherd",
    }

	printInfo(dog)

	gorilla := Gorilla{
        name:          "Jock",
        color:         "grey",
        numberOfTeeth: 38,
    }

	printInfo(gorilla)
	

}

func printInfo(b Animal) {
	fmt.Println("This animal says", b.says(), "and has", b.numberOfLegs(), "legs")
}

func (d Dog) says() string {
	return "woof"
}

func (d Dog) numberOfLegs() int {
	
	return 14
}

func (d Gorilla) says() string {
	return "puff"

}

func (d Gorilla) numberOfLegs() int {
	
	return 14
}
