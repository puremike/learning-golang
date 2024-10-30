package main

import (
	"fmt"

	"github.com/puremike/09package/helpers"
)

func main() {
	id := helpers.User{
		Name:     "John Doe",
        Occupation: "Software Developer",
        Age:       30,
	}

	fmt.Println(id)

}
