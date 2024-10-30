package main

import (
	"log"
	"severalfunc/helpers"
)

func main() {
	myChannel := make(chan int)
	defer close(myChannel)
	go useGeneratedRandomNFunc(myChannel, 100)
	num := <- myChannel

	log.Println("Random Number =", num)
}

func useGeneratedRandomNFunc(n chan int, m int) {
	n <- helpers.GenerateRandomNumbers(m)
}