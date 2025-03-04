package main

import (
	"fmt"
	"time"
)

func useRoutines() {

	fmt.Printf("\nUsing Routines...........\n")
	go makeLatte()
	makeBlackCoffee()
}

func makeLatte() {
	fmt.Println("Making a latte")
	time.Sleep(2 * time.Second)
	fmt.Println("Latte is ready")
}

func makeBlackCoffee() {
	fmt.Println("Making a black coffee")
	time.Sleep(1 * time.Second)
	fmt.Println("Black coffee is ready")
}

// The go keyword lets makeLatte() run in the background while makeBlackCoffee() happens at the same time. Without go, the latte would finish first, then the black coffee would start.
