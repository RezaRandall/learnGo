package main

import (
	"fmt"
	"strings"
)

func main() {
	// PENERAPAN METHOD
	var s1 = student{"Reza Randall", 30}
	s1.sayHello()

	var name = s1.getNameAt(2)
	fmt.Println("nama panggilan : ", name)
	fmt.Println()
	// METHOD POINTER
	fmt.Println("s1 before ", s1.name) // Reza Randall

	s1.changeNameOne("Jason Bourne")
	fmt.Println("s1 after changeName1", s1.name)

	s1.changeNameTwo("Bungut")
	fmt.Println("s1 after changeName2", s1.name)
	fmt.Println()

	// PENGAKSESAN METHOD DARI VARIABLE OBJECT BIASA
	var S1 = student{"Ayu Irmawati", 44}
	S1.sayHello()

	// PENGAKSESAN METHOD DARI OBJECT POINTER
	S2 := &student{"Real Ika", 45}
	S2.sayHello()
}

// PENERAPAN METHOD
type student struct {
	name  string
	grade int
}

func (s student) sayHello() {
	fmt.Println("hello", s.name)
}

func (s student) getNameAt(i int) string {
	return strings.Split(s.name, " ")[i-1]
}

// METHOD POINTER
func (s student) changeNameOne(name string) {
	fmt.Println("---> On changeName1, name change to ", name)
	s.name = name
}

func (s student) changeNameTwo(name string) {
	fmt.Println("---> On changeName2, name change to ", name)
	s.name = name
}
