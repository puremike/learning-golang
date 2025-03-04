// Concurrency is the ability to handle multiple tasks at once, making progress on them without needing to finish one before starting another. Go routines and channels make concurrency easy in Go.

package main

import "fmt"

func orderChannel(order string, ch chan string) {
	fmt.Println("Order received:", order)
	ch <- order + " is ready"
}

func orderMain() {
	fmt.Printf("\nUsing Concurrency...........\n")
	orderMessage := make(chan string)
	go orderChannel("Latte", orderMessage)
	go orderChannel("Black Coffee", orderMessage)

	fmt.Println("Order 1:", <-orderMessage)
	fmt.Println("Order 2:", <-orderMessage)
}
