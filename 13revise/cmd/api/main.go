package main

import "fmt"

func main() {
	useIt, useIt2 := revisingVariable()
	fmt.Println(useIt, useIt2)

	usePointer()
	useTypeStruct()
	structFunc()
	mapSlices()
	fmt.Println("Factorial = ", factorialN(4))
	fmt.Println("Factorial = ", factorialN(5))
	useRoutines()
	usePastry()
	orderMain()
}
