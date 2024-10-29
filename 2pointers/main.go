package main

import (
	"fmt"
	"log"
)

func main() {
	var color = "White"

	log.Println("My favorite color is", color)
	log.Println("Color before calling the changeColorWithPointer function is", color)

	changeColorWithPointer(&color)
	log.Println("Color after calling the changeColorWithPointer function is", color)

	whatWeDid := changeColorWithPointer(&color)
	log.Println("Here is what we did:", whatWeDid)

	fmt.Printf("\n")
	understandPointer(42, 2701)
}

func changeColorWithPointer(c *string) string {
	newColor := "Black"
	*c = newColor   // Change the color to "Black"

	return "We changed the color using pointer to " + *c
}

func understandPointer(i, j int) {

	fmt.Println("I =", i)
	p := &i // point p to i
	fmt.Println("P =", p) // print the memory location of p and not what it points to
	fmt.Println("P =", *p) 

	*p = 21 // change what p points to - update the value of i
	fmt.Println("New Value of I =", i) 
	fmt.Println("New Value of I =", *p)

	fmt.Println("J =", j)
	r := &j
	fmt.Println("R =", *r)
	*r	= *r / 37
	fmt.Println("J =", *r)

}

