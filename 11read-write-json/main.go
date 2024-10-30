package main

import (
	"encoding/json"
	"fmt"
)

type Config struct {
	Name       string `json:"name"`
	Age        int    `json:"age"`
	City       string `json:"city"`
	Occupation string `json:"occupation"`
}

func main() {

	miJSON := `[
		{
			"name": "John Doe",
			"age": 30,
			"city": "New York",
			"occupation": "Software Engineer"
		},
		{
			"name": "Jane Doe",
			"age": 25,
			"city": "Los Angeles",
			"occupation": "Product Manager"
		}
	]`

	var users []Config

	err := json.Unmarshal([]byte(miJSON), &users)
	if err != nil {
		fmt.Println("Error unmarshalling data:", err)
		return
	}

	fmt.Println("Unmarshalled Data:", users)

	// WRITE FROM STRUCT TO JSON

	fmt.Println("......WRITE FROM STRUCT TO JSON.......")

	var myJSON []Config

	firstData := Config {
		Name: "Jackson Doe",
		Age: 25,
		City: "California",
		Occupation: "Software Manager",
	}

	myJSON = append(myJSON, firstData)

	secondData := Config{

		Name: "John Doe",
		Age: 85,
		City: "California",
		Occupation: "Golang Developer",
	
	}

	myJSON = append(myJSON, secondData)


	newJSON, err := json.MarshalIndent(myJSON, "", "    ")

	if err != nil {
		fmt.Println("Error marshalling data", err)
	}

	fmt.Println("Marshalled Data : ", string(newJSON))
}
