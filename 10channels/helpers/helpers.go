package helpers

import "math/rand"

func GenerateRandomNumbers(ch chan int, n int) {
	num := rand.Intn(n)
	ch <- num // pass the message (num) to the channel
}
