package main

// factorialN returns the factorial of a given integer n.
// If n is less than 1, the function returns n.

func factorialN(n int) int {

	if n < 1 {
		return n
	}

	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}
	return fact

}
