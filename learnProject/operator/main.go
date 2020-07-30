package main

import "fmt"

func main() {
	// Operator terdapat 3 kategori operator secara umum : ariitmatika, perbandingan dan logika

	// Operator aritmatika
	var value = (((2+6)%3)*4 - 2) / 3
	fmt.Println(value)

	// Operator Perbandingan
	var isEqual = (value == 2)                     // hasil boolean
	fmt.Printf("Nilai %d (%t) \n", value, isEqual) // Untuk memunculkan nilai bool menggunakan fmt.Printf(), bisa gunakan layout format %t

	// Operator logika
	var left = false
	var right = true

	var leftAndRight = left && right                    // kiri dan kanan
	fmt.Printf("left and right \t(%t)\n", leftAndRight) // leftAndRight bernilai false, karena hasil dari false dan true adalah false.

	var leftOrRight = left || right                     // kiri atau kanan
	fmt.Printf("left or right \t (%t) \n", leftOrRight) // leftOrRight bernilai true, karena hasil dari false atau true adalah true.

	var leftReserve = !left                       // negasi atau nilai kebalikan
	fmt.Printf("!left \t\t (%t) \n", leftReserve) // leftReverse bernilai true, karena negasi (atau lawan dari) false adalah true.
	// Template \t digunakan untuk menambahkan indent tabulasi. Biasa dimanfaatkan untuk merapikan tampilan output pada console.
}
