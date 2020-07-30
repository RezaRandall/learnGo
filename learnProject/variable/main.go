package main

import "fmt"

func main() {
	var firstName string = "Daniel"
	var lastName string
	lastName = "Rubangka"
	fmt.Printf("Haloo "+"%s %s \n", firstName, lastName+" !!")

	nickName := `luna` // inference type deklarasi variable type datanya berdasarkan nailainya
	age := 30
	fmt.Println("My nick name is", nickName+" I'am", age, " Years old")

	midName := "lulu"
	midName = "lala"

	fmt.Println(midName)

	// DEKLARASI MULTI VARIABLE
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"
	fmt.Println(first, second, third)

	fourth, fifth, sixth := "empat", 5, "enam"
	one, isSaturday, fourPointOne, say := "one", true, 4.2, "haii"
	fmt.Println(fourth, fifth, sixth)
	fmt.Println(one, isSaturday, fourPointOne, say)

	/*	VARIABLE UNDERSCORE _
		reserve variable yang bisa dimanfaatkan untuk menampung nilai yang tidak dipakai. Bisa dibilang variabel ini merupakan keranjang sampah.
	*/
	_ = "Belajar Golang  itu Mudah"
	_ = "Belajar Golang itu Menyenangkan"

	nama, _ := "john", "wick"
	fmt.Println(nama)

	// DEKLARASI VARIABLE MENGGUNAKAN KEYWORD new
	newName := new(string)

	/* Variabel name menampung data bertipe pointer string.
	Jika ditampilkan yang muncul bukanlah nilainya melainkan alamat memori nilai tersebut */
	fmt.Println(newName) // 0xc00002e2b0

	/* Untuk menampilkan nilai aslinya, variabel tersebut perlu di-dereference terlebih dahulu, menggunakan tanda asterisk (*). */
	fmt.Println(*newName) // ""

}
