package main

import "fmt"

func main() {
	var fruits = []string{"Apple", "Grape", "Banana", "Orange"}
	fmt.Println(fruits)
	fmt.Println(fruits[0])

	// SALAH SATU PERBEDAAN SLICE DAN ARRAY DAPAT DIKETAHUI PADA SAAT DEKLARASI VARIABLENYA, JIKA JUMLAH ELEMENNYA TIDAK DITULISKAN
	// MAKA VARIABLE TERSEBUT ADALAH SLICE

	// HUBUNGAN SLICE DENGAN ARRAY & OPERASI SLICE
	var newFruits = fruits[0:4]
	// Kode fruits[0:2] maksudnya adalah pengaksesan elemen dalam slice fruits yang dimulai dari indeks ke-0, hingga elemen sebelum indeks ke-2
	// 0 adalah index mulai AMBIL DARI & 3 adalah index BATAS DARI atau sebelum index ke-3
	fmt.Println(newFruits)

	fmt.Println(fruits[0:2]) // semua elemen mulai indeks ke-0, hingga sebelum indeks ke-2
	fmt.Println(fruits[0:4]) // semua elemen mulai indeks ke-0, hingga sebelum indeks ke-4
	fmt.Println(fruits[0:0]) // menghasilkan slice kosong, karena tidak ada elemen sebelum indeks ke-0
	fmt.Println(fruits[4:4]) // menghasilkan slice kosong, karena tidak ada elemen yang dimulai dari indeks ke-4
	// fmt.Println(fruits[4:0]) // error, pada penulisan fruits[a,b] nilai a harus lebih besar atau sama dengan b
	fmt.Println(fruits[:])        // semua elemen
	fmt.Println(fruits[2:])       // semua elemen mulai indeks ke-2
	fmt.Print(fruits[:2], "\n\n") // semua elemen hingga sebelum indeks ke-2

	// SLICE MERUPAKAN TYPE DATA REFERENCE
	var aFruits = fruits[0:3]
	var bfruits = fruits[1:4]
	var aaFruits = aFruits[1:2]
	var bbFruits = bfruits[0:1]

	fmt.Println(fruits)
	fmt.Println(aFruits)
	fmt.Println(bfruits)
	fmt.Println(aaFruits)
	fmt.Print(bbFruits, "\n\n")

	fmt.Println("BUAH GRAPE DIRUBAH MENJADI DURIAN")
	bbFruits[0] = "Durian"

	fmt.Println(fruits)
	fmt.Println(aFruits)
	fmt.Println(bfruits)
	fmt.Println(aaFruits)
	fmt.Print(bbFruits, "\n\n")

	// FUNGSI len()
	/*
		Fungsi len() digunakan untuk menghitung jumlah elemen slice yang ada.
		Sebagai contoh jika sebuah variabel adalah slice dengan data 4 buah,
		maka fungsi ini pada variabel tersebut akan mengembalikan angka 4.
	*/
	fmt.Print(len(fruits), "\n\n")

	// FUNGSI cap()
	/*
		Fungsi cap() digunakan untuk menghitung lebar atau kapasitas maksimum slice.
		Nilai kembalian fungsi ini untuk slice yang baru dibuat pasti sama dengan len,
		tapi bisa berubah seiring operasi slice yang dilakukan. Agar lebih jelas, silakan disimak kode berikut
	*/
	fmt.Println(len(fruits), "len fruits")
	fmt.Println(cap(fruits), "cap fruits")

	fmt.Println(len(aFruits), "len aFruits")
	fmt.Println(cap(aFruits), "cap aFruits")
	fmt.Println(len(bfruits), "len bFruits")
	fmt.Print(cap(bfruits), " cap bFruits \n\n")

	// FUNGSI append()
	var cFruit = append(fruits, "Strawberry")

	fmt.Println(fruits)
	fmt.Println(cFruit, "\n")

	var haub = []string{"Apple", "Grape", "Banana"}
	var bHaub = haub[0:2]

	fmt.Println(cap(bHaub), "Capacity is cap of bHaub")
	fmt.Println(len(bHaub), "lenght is len of bHaub")

	fmt.Println(haub)
	fmt.Println(bHaub, "\n")

	var cHaub = append(bHaub, "Kelopo")
	var dHaub = append(cHaub, "Rambutan")
	fmt.Println(haub)
	fmt.Println(bHaub)
	fmt.Println(cHaub)
	fmt.Println(dHaub, "\n")

	// FUNGSI copy()
	var dst = make([]string, 3)
	src := []string{"kelepon", "pastel", "nugget", "cenil"}

	var n = copy(dst, src) // di copy untuk dst, diambil dari src

	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println(n)

	dataSatu := []string{"Melon", "Coconut", "Manggo"}
	dataDua := []string{"Coffee", "Beer"}

	cpDtSatuDanDua := copy(dataSatu, dataDua)
	fmt.Println(dataSatu)
	fmt.Println(dataDua)
	fmt.Println(cpDtSatuDanDua, "\n")

	// PENGAKSESAN ELEMENT SLICE DENGAN 3 INDEKS
	var aDataSatu = dataSatu[0:2]
	var bDataSatu = dataSatu[0:2:2]

	fmt.Println(dataSatu)
	fmt.Println(len(dataSatu), "len dari dataSatu")
	fmt.Println(cap(dataSatu), "cap dari data satu\n")

	fmt.Println(aDataSatu)
	fmt.Println(len(aDataSatu))
	fmt.Println(cap(aDataSatu), "\n")

	fmt.Println(bDataSatu)
	fmt.Println(len(bDataSatu))
	fmt.Println(cap(bDataSatu), "\n")

	// NEXT A.16 MAP
}
