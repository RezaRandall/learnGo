package main

import (
	f "fmt"
	"levelAccess/library"
)

func main() {
	// library.SayHello("Ethan")
	// library.introduce("Evan")

	var s1 = library.Student{"Reza Irawan", 30}
	f.Println("nama : ", s1.Name)
	f.Println("Grade : ", s1.Grade)

	SayHello("nana")
}
