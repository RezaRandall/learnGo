package main

import "fmt"

func main() {
	fmt.Println("Hello World", "my name is joker", "hahaa")

	var firstName string = "Reza"

	var lastName string
	lastName = "Randall"

	fmt.Printf("halo %s %s \n\n", firstName, lastName)
	fmt.Println("halo", firstName, lastName+"!")

	// Deklarasi variable tanpa type data
	nickName := "Robert"

	fmt.Printf("Hey %s %s \n \n", firstName, nickName)

	// Mendeklarasikan multple variable

	var first, second, third string
	first, second, third = "lolo", "lala", "lili"

	fmt.Println(first, second, third)

	var fourth, fifth, sixth string = "empat", "lima", "enam"
	seventh, eight, ninth := "tujuh", "delapan", "sembilan"

	fmt.Println(fourth, fifth, sixth, seventh, eight, ninth)

	one, isFriday, twoPointOne, say := 1, true, 2.1, "hai"

	fmt.Println(one, isFriday, twoPointOne, say)

	// Variable Underscore _
	_ = "Belajar Golang"
	_ = "Golang itu mudah"
	nama, _ := "rere", "raaa"

	fmt.Println(nama, `\n`)

	// Deklarasi variable menggunakan keyword new
	name := new(string)

	fmt.Println(name)
	fmt.Println(*name)

	// @@ BERAKHIR DI A.8 VARIABLE, NEXT TYPE DATA @@
}
