package main

import "fmt"

func anotherFunc() (string, string){
	return "This is another function", "that works."
}

func anotherWayToDeclareVar() {
	var i, j int = 1, 2
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

func main() {
	fmt.Println("Hey, Michael. Welcome to Golang!")

	var firstStatement string; var firstNumber int
	firstStatement = "Hello, World!"
	firstNumber = 42

	secondNumber := firstNumber + 1

	fmt.Println("His first statement and first number are ", firstStatement, " and ", firstNumber, " respectively.")
	fmt.Println("His favorite number is ", secondNumber)

	// var holdMySecondFunc string; var holdMySecondFunck string
	// holdMySecondFunc, holdMySecondFunck = anotherFunc()

	holdFunc, holdFunck := anotherFunc()

	fmt.Println(holdFunc, holdFunck)

	fmt.Println("\nAnother way to declare variable")
	anotherWayToDeclareVar()
}