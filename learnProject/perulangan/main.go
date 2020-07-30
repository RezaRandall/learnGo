package main

import "fmt"

func main() {
	var i = 0

	// PERULANGAN MENGGUNAKAN KEYWORD for
	// for i = 0; i < 5; i++ {
	// 	fmt.Println("Angka", i)
	// }

	// PENGGUNAAN KEYWORD for DENGAN ARGUMENT HANYA KONDISI
	// for i < 5 {
	// 	fmt.Println("Nilai", i)
	// 	i++
	// }

	// PENGGUNAAN KEYWORD for TANPA ARGUMENT
	// for {
	// 	fmt.Println("Hasill", i)
	// 	i++

	// 	if i == 6 {
	// 		break
	// 	}
	// }

	// PENGGUNAAN KEYWORD BREAK & CONTINUE
	// for i = 1; i < 12; i++ {
	// 	if i%2 == 1 {
	// 		continue
	// 	}

	// 	if i > 8 {
	// 		break
	// 	}
	// 	fmt.Println("Hasil", i)
	// }

	// PERULANGAN BERSARANG
	/*
		Pada kode di bawah, untuk pertama kalinya fungsi fmt.Println() dipanggil tanpa disisipkan parameter.
		Cara seperti ini bisa digunakan untuk menampilkan baris baru. Kegunaannya sama seperti output dari statement fmt.Print("\n")
	*/

	for i = 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}

	// PEMANFAATAN LABEL DALAM PERULANGAN
outerloop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				break outerloop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}

}
