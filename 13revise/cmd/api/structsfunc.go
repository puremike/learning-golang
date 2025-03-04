package main

import "fmt"

type User struct {
	name string
}

func (u *User) useStructWithFunc() string {
	return u.name
}

func structFunc() {
	user1 := User{
		name: "Foo",
	}

	user2 := User{
		name: "Bar",
	}

	fmt.Printf("\nUsing Structs With Functions.........\n")

	fmt.Printf("My name is %s\n", user1.useStructWithFunc())
	fmt.Printf("My name is %s\n", user2.useStructWithFunc())
}
