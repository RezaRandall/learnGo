package main

import "fmt"

func main() {
	// A.12 SELEKSI KONDISI
	point := 8
	if point == 10 {
		fmt.Println(`lulus dengan nilai sempurna`)
	} else if point > 5 {
		fmt.Println(`lulus`)
	} else if point == 4 {
		fmt.Println(`hampir lulus`)
	} else {
		fmt.Printf("tidak lulus, nilai anda %d \n", point)
	}

	nilai := 90
	if nilai < 10 {
		fmt.Println("anda lolos dengan selamat")
	} else {
		fmt.Println("anda terlindas")
	}

	// variable temporary pada if else
	points := 8840.20

	if percent := points / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect !\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good !\n", percent, "%")
	} else {
		fmt.Printf("%1f%s not bad !\n", percent, "%")
	}

	// seleksi kondisi menggunakan keyword switch - case
	// Switch merupakan seleksi kondisi yang sifatnya fokus pada satu variabel,
	// lalu kemudian di-cek nilainya. Contoh sederhananya seperti penentuan apakah nilai variabel x adalah: 1, 2, 3, atau lain

	var testPoint = 6

	switch testPoint {
	case 8:
		fmt.Println(`awesome`)
	case 7:
		fmt.Println("perfect")
	default:
		fmt.Println(`not bad`)
	}

	// PEMANFAATAN BANYAK case UNTUK BANYAK KONDISI

	switch testPoint {
	case 8:
		fmt.Println("awesome")
	case 7, 6, 5, 4:
		fmt.Println("perfect")
	default:
		fmt.Println("not bad")
	}

	// KURUNG KURAWAL PADA KEYWORD case & default
	switch testPoint {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can try again")
		}
	}

	// Switch dengan gaya if - else
	switch {
	case testPoint == 8:
		fmt.Println("Perfect")
	case (testPoint < 8) && (point > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need learn more")
			fmt.Print("\n")
		}
	}

	// Penggunaan keyword fallthrought Dalam switch
	/*
		Keyword fallthrough digunakan untuk memaksa proses pengecekkan diteruskan ke case selanjutnya dengan
		TANPA MENGHIRAUKAN NILAI KONDISINYA, jadi case di pengecekan selanjutnya tersebut selalu dianggap benar
		(meskipun aslinya adalah salah)
	*/
	switch {
	case testPoint == 8:
		fmt.Println("perfect")
	case (testPoint < 8) && (testPoint > 3):
		fmt.Println("awesome")
		fallthrough
	case testPoint < 5:
		fmt.Println(`not bad`)
	default:
		fmt.Println(`not bad`)
		fmt.Println("you need learn more")
	}

	// Seleksi kondisi bersarang
	/*
		Seleksi kondisi bersarang adalah seleksi kondisi, yang berada dalam seleksi kondisi,
		yang mungkin juga berada dalam seleksi kondisi,
		dan seterusnya. Seleksi kondisi bersarang bisa dilakukan pada if - else, switch, ataupun kombinasi keduanya.
	*/

	pointOfView := 0
	if pointOfView > 7 {
		switch pointOfView {
		case 10:
			fmt.Println("Perfect")
		default:
			fmt.Println("nice")
		}
	} else {
		if pointOfView == 5 {
			fmt.Println(`not bad`)
		} else if pointOfView == 3 {
			fmt.Println(`keep trying`)
		} else {
			fmt.Print("you can do it ")
			if pointOfView == 0 {
				fmt.Println("TRY HARDER!!")
			}
		}
	}
}
