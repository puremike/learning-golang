package main

import "fmt"

func usePointer() {

	fmt.Printf("\nUsing Pointers...........\n")
	color := "Blue"
	fmt.Println("Color before calling the pointerOne function is", color)

	newC := pointerOne(&color)
	fmt.Println("Color after calling the pointerOne function is", newC)

	pointerTwo(42, 2701)
}

func pointerOne(c *string) string {
	newColor := "Black"
	*c = newColor // Change the color to "Black"
	return *c
}

func pointerTwo(i, j int) {

	fmt.Println("I =", i)
	p := &i               // point p to i
	fmt.Println("P =", p) // print the memory location of p and not what it points to
	fmt.Println("P =", *p)

	*p = 21 // change what p points to - update the value of i
	fmt.Println("New Value of I =", i)
	fmt.Println("New Value of I =", *p)

	fmt.Println("J =", j)
	r := &j
	fmt.Println("R =", *r)
	*r = *r / 37
	fmt.Println("J =", *r)
}
