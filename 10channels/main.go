package main

import (
	"log"
	"severalfunc/helpers"
)

func main() {
	myChannel := make(chan int) // create a channel
	defer close(myChannel)

	go helpers.GenerateRandomNumbers(myChannel, 100)

	num := <-myChannel
	log.Println("Random Number =", num)
}
