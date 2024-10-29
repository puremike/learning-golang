package main

import "log"

type myStruct struct {
	firstName string
}

func (s *myStruct) accessStructPointer() string {
	return s.firstName
}

func main() {

	var myVar1 myStruct
	myVar1.firstName = "foo"

	myVar2 := myStruct{
		firstName: "bar",
	}

	// log.Println(myVar1.firstName, myVar2.firstName)
	log.Println(myVar1.accessStructPointer(), myVar2.accessStructPointer())

}