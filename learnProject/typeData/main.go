package main

import "fmt"

func main() {
	var positiveNumber uint8 = 89
	var negativeNumber = -1234567
	fmt.Printf("Bilangan positiv: %d \n", positiveNumber)
	fmt.Printf("Bilangan negative: %d \n", negativeNumber) // %d yg BERARTI DECIMAL DIGUNAKAN UNTUK MEMFORMAT DATA DECIMAL (NON DECIMAL)

	// TYPE DATA NUMERIK DECIMAL ATAU KOMA(,) ATAU FLOAT
	decimalNumber := 2.89

	fmt.Printf("Bilangan Decimal: %f\n", decimalNumber)
	fmt.Printf("Bilangan Decimal: %.3f\n", decimalNumber)

	// TYPE DATA bool (BOOLEAN) true & false
	var exsist bool = true
	fmt.Printf("Exsist? %t \n", exsist) // %t DIGUNAKAN UNTUK MEMFORMAT DATA bool (BOOLEAN)

	// TYPE DATA STRING
	var message string = "Apa kabar dunia?"
	fmt.Printf("Message: %s \n", message)

	var pesan string = `Nama saya "Dwi Reza Irawan".
Salam Kenal
Mari kita belajar "Golang".`
	fmt.Println(pesan)

	/* NILAI nil & ZERO VALUE
	- nil benar benar kosong
	- Zero Value string ""
	- Zero Value bool false
	- Zero Value int 0
	- Zero Value float 0.0 	*/
}
