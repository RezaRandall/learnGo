package main

import "fmt"

func main() {
	var names [4]string

	names[0] = " Dwi "
	names[1] = " Reza "
	names[2] = " Irawan "
	names[3] = " Randall "
	fmt.Println(names[0], names[1], names[2], names[3])

	// CARA HORIZONTAL 
	var fruits = [4]string{"apple", "manggo", "banana", "mengkudu"}

	fmt.Println("Jumlah element \t\t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)
	/*
		Penggunaan fungsi fmt.Println() pada data array tanpa mengakses indeks tertentu, 
		akan menghasilkan output dalam bentuk string dari semua array yang ada. 
		Teknik ini biasa digunakan untuk debugging data array.
	*/

	// INISIALISASI ARRAY DENGAN GAYA VERTICAL
	var cars = [4]string{
		"lamborghini", 
		"Ferarry", 
		"Mazda", 
		"Audi",
		/*
			Khusus untuk deklarasi array dengan cara vertikal, tanda koma wajib dituliskan setelah elemen, 
			termasuk elemen terakhir. Jika tidak, maka akan muncul error.
		*/
	}
	fmt.Println(cars)

	// INISIALISASI NILAI AWAL ARRAY TANPA JUMLAH ELEMENT
	var motoBike = [...]string{"yahama", "honda", "ducati"}
	var numbers = [...]int{2,3,4,53,3,65,6,4}

	fmt.Println("data array motor bike: \t", motoBike)
	fmt.Println("Jumlah Array Motor Bike: \t", len(motoBike))
	fmt.Println("Data Array Numbers: \t", numbers)
	fmt.Println("Jumlah Array Numbers: \t", len(numbers))
	
	// ARRAY MULTI DIMENSI 
	var angka1 = [2][3]int{[3]int{2,3,4}, [3]int{4,5,6}}
	var angka2 = [2][3]int{{4,3,2}, {7,6,5}}

	fmt.Println("Angka1 :", angka1)
	fmt.Println("Angka2 :", angka2)

	// PERULANGAN ARRAY MENGGUNAKAN KEYWORD for
	for i := 0; i < len(motoBike); i++{
		fmt.Printf("element %d : %s\n", i, motoBike[i])
	}

	fmt.Println()
	// PERULANGAN ELEMENT Array MENGGUNAKAN KEYWORD for - range
	// array (motoBike) diambil elemennya secara berurutan, nilai setiap element ditampung variable oleh (bike)
	// sedangkan indexnya ditampung oleh index i 
	for i, bike := range motoBike{
		fmt.Printf("element %d : %s \n", i, bike)
	}
	
	// PENGGUNAAN VARIABLE UNDERSCORE _ DALAM for - range 
	for _, bike := range motoBike{
		fmt.Printf("nama sepeda motor : %s\n", bike)
	} 
	// PADA KODE DIATAS, YG SEBELUMNYA ADALAH VARIABLE i DIGANTI DENGAN  _, KARENA KEBETULAN VARIABLE i TIDAK DIGUNAKAN 
	// JIKA YG DIBUTUHKAN HANYA INDEXNYA SAJA, BISA DIGUNAKAN 1 BUAH VARIABLE SAJA SETELAH KEYWORD for 
	/*
		for i, _ := range motoBike 
		atau 
		for i := range motoBike 
	*/
	for i := range motoBike{
		fmt.Printf("index : %d \n", i)
	}

	for _, bikes  := range motoBike{
		fmt.Printf("element Bike : %s \n", bikes)
	}

	// ALOOKASI ELEMENT ARRAY MENGGUNAKAN KEYWORD make
	// deklarasi sekaligus alokasi data array juga bisa dilakukan lewat keyword make
	
	// PARAMETER KEYWORD PERTAMA make DIISI DENGAN TIPE DATA ELEMENT ARRAY YG DI INGINKAN 
	// PARAMETER KEDUA ADALAH JJUMLAH ELEMENTNYA. PADA KODE DIBAWAH VARIABLE buahBuah TERCETAK SEBAGAI ARRAY STRING DENGAN 2 ALOKASI 2 SLOT
	var buahBuah = make([]string, 2)
	buahBuah[0] = "Anggur"
	buahBuah[1] = "Nangka"

	fmt.Println(buahBuah)


}
