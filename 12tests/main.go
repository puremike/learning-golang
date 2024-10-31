package main

import (
	"errors"
	"fmt"
)

func main() {

	var a,b float64 = 10.0, 20.0

	fmt.Println("A function that performs division operation")
	result, err := divide(a,b)

	if err != nil {
		fmt.Println("Failed to divide")
	}

	fmt.Printf("%.2f / %.2f = %.2f", a, b, result)

}

func divide(a, b float64) (float64, error) {
	result := a/b

	if b == 0 {
		return result, errors.New("error: cannot divide by zero")
	}

	return result, nil
}

