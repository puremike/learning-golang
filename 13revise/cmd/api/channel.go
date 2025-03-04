// A Go channel is like a walkie-talkie or a conveyor belt that lets Go routines talk to each other or pass things (data) between them safely.

package main

import "fmt"

func makePastry(ch chan string) {
	fmt.Println("Making Pastry!")
	ch <- "Pastry is ready" // Send message through the channel
}

func usePastry() {
	fmt.Printf("\nUsing Channels...........\n")

	pastryMessage := make(chan string) // create a channel
	go makePastry(pastryMessage)       // Start Go routine to make pastry

	// Wait for the pastry (receive from channel)
	message := <-pastryMessage
	fmt.Println(message)

	fmt.Println("Serving latte with pastry")
}


