package helpers

import "math/rand"

func GenerateRandomNumbers(n int) int {
	num := rand.Intn(n)
	return num
} 