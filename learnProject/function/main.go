package main

import (
	"fmt"
	"math/rand" // Fungsi rand.Seed() berada dalam package math/rand, yang harus di-import terlebih dahulu sebelum bisa dimanfaatkan.
	"strings"
	"time" // Package time juga perlu di-import karena kita menggunakan fungsi (time.Now().Unix()) disitu
)

// PENERAPAN FUNGSI
func main() {
	var names = []string{"Reza", "Irawan"}
	printMessage("Hallo", names)

	// FUNGSI DENGAN RETURN VALUE / NILAI BALIK
	rand.Seed(time.Now().Unix())
	/*
		Fungsi ini diperlukan untuk memastikan bahwa angka random yang akan di-generate benar-benar acak.
		Kita bisa gunakan angka apa saja sebagai nilai parameter fungsi ini (umumnya diisi time.Now().Unix()).
		Fungsi rand.Seed() berada dalam package math/rand, yang harus di-import terlebih dahulu sebelum bisa dimanfaatkan.
		Package time juga perlu di-import karena kita menggunakan fungsi (time.Now().Unix()) disitu
	*/
	var randomValue int

	randomValue = randomWithRange(2, 10)
	fmt.Println("Random Number : ", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("Random Number : ", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("Random Number : ", randomValue)

	// Penggunaan keyword return untuk menghentikan proses dalam fungsi
	devideNumber(10, 2)
	devideNumber(6, 0)
	devideNumber(8, -2)
}

func devideNumber(m, n int) {
	if n == 0 || m == 0 {
		fmt.Printf("Invalid Devider, %d cannot divide by %d \n", m, n)
		return
	}

	var res = m / n
	fmt.Printf("%d / %d = %d\n", m, n, res)
}

func randomWithRange(min, max int) int {
	var value = rand.Int()%(max-min+1) + min
	return value
}

func printMessage(message string, arrName []string) {
	var nameString = strings.Join(arrName, " ")
	fmt.Println(message, nameString)

}
