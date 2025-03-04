package main

import "fmt"

type User1 struct {
	name string
}

func mapSlices() {

	fmt.Printf("\nUsing Maps.........\n")

	mapOne := make(map[string]string)
	mapOne["name"] = "FooBar"

	mapTwo := make(map[string]int)
	mapTwo["age"] = 30
	mapTwo["one-hundred"] = 100

	mapThree := make(map[string]User1)
	mapThree["name"] = User1{
		name: "John",
	}
	mapThree["lastname"] = User1{
		name: "Doe",
	}

	fmt.Printf("Name: %s\n Age: %d\n Fav Number: %d\n", mapOne["name"], mapTwo["age"], mapTwo["one-hundred"])
	fmt.Printf("Name: %s\n LastName: %s\n", mapThree["name"].name, mapThree["lastname"].name)

	fmt.Printf("\nUsing Slices.........\n")
	color := []string{"White", "Black", "Red", "Blue"} // a dynamic array
	fmt.Printf("Favorite Colors: %s\n", color)
	number := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // static array of 10
	fmt.Printf("Favorite Numbers: %d\n", number)
	color = append(color, "Green", "Yellow") // added more colors to the slice
	fmt.Printf("Favorite Colors: %s\n", color)

	fmt.Println("Another way to declare a slice .... here, the array has to be dynamic")
	var colorTwo []string
	colorTwo = append(colorTwo, "Black", "White", "Purple", "Gold")

	var numberTwo []int
	numberTwo = append(numberTwo, 1, 2, 3, 4, 5, 6, 7)

	fmt.Printf("Favorite Colors: %s\n", colorTwo)
	fmt.Printf("Favorite Numbers: %d\n", numberTwo)

	fmt.Println("Length of ColorTwo:", len(colorTwo))
	fmt.Println("Length of NumberTwo:", len(numberTwo))
}
