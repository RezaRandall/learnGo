package main

import (
	"fmt"
	"strings"
)

func main() {
	// PENERAPAN FUNGSI VARIADIC
	var avg = calculate(2, 4, 5, 4, 3, 2, 4, 5, 5, 6, 5, 4, 5, 4, 3, 4, 2, 4, 3, 5)

	/*
		Fungsi fmt.Sprintf() pada dasarnya sama dengan fmt.Printf(), hanya saja fungsi ini tidak menampilkan nilai,
		melainkan mengembalikan nilainya dalam bentuk string. Pada kasus di atas, nilai kembalian fmt.Sprintf() ditampung oleh variabel msg.
		Selain fmt.Sprintf(), ada juga fmt.Sprint() dan fmt.Sprintln().
	*/
	var msg = fmt.Sprintf("Rata - rata nilai :  %2.f", avg)
	fmt.Println(msg)

	// PENGISIAN PARAMETER FUNGSI VARIADIC MENGGUNAKAN DATA SLICE
	var nomors = []int{2, 3, 4, 2, 4, 3, 1, 2, 3, 4, 4, 4, 5, 4, 3, 3, 3, 2, 3, 1, 2, 3, 3, 2, 11, 1}
	var avr = calculate(nomors...)
	var pesan = fmt.Sprintf("Rata - rata \t :  %2.f", avr)

	fmt.Println(pesan)

	yourHobbys("Reza", "Running", "Skateboarding", "Studying")

}

func calculate(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	// Sebelumnya sudah dibahas bahwa float64 merupakan tipe data. Tipe data jika ditulis sebagai fungsi (penandanya ada tanda kurungnya)
	// berguna untuk casting. Casting sendiri adalah teknik untuk konversi tipe sebuah data ke tipe lain.
	var avg = float64(total) / float64(len(numbers))
	return avg
}

// PENGGABUNGAN FUNGSI DENGAN PARAMETER BIASA & VARIADIC
// Parameter variadic bisa dikombinasikan dengan parameter biasa, dengan syarat parameter variadic-nya harus diposisikan di akhir.
// Contohnya bisa dilihat pada kode berikut.

// Nilai parameter pertama fungsi yourHobbies() akan ditampung oleh name,
// sedangkan nilai parameter kedua dan seterusnya akan ditampung oleh hobbies sebagai slice.
func yourHobbys(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")

	fmt.Printf("My name is %s \n", name)
	fmt.Printf("My hobby's are %s \n", hobbiesAsString)
}
