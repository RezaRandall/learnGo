package main

import "fmt"

// create new folder mkdir <folder name>
// delete folder rmdir /s <folder name>
// create new file type nul > <file name .extention>
// delete file del <file name .extension>
// go mod init <modul name>

func main() {
	// A.10 KONSTANTA
	// Konstanta adalah jenis variabel yang nilainya tidak bisa diubah.
	// Inisialisasi nilai hanya dilakukan sekali di awal, setelah itu variabel tidak bisa diubah nilainya.
	const namaDepan string = "jhon"
	fmt.Print("Halo ", namaDepan, "!\n")

	// teknik inference yaitu cukup dengan menghilangkan tipr data pada saat deklarasi
	const namaAkhir = "Winchester"
	fmt.Print("Nice to meet you ", namaAkhir, "!\n")

	fmt.Println("john wick")
	fmt.Println("john", "wick")

	fmt.Print("john wick\n")
	fmt.Print("john ", "wick\n")
	fmt.Print("john", " ", "wick\n")
}
