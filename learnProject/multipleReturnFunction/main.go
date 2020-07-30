package main

import (
	"fmt"
	"math"
)

// PENERAPAN FUNGSI MULTIPLE RETURN
func main() {
	var diameter float64 = 15
	var area, circumference = calculate(diameter)

	fmt.Printf("Luas lingkaran \t\t : %.2f \n", area)
	fmt.Printf("Keliling lingkaran \t : %.2f \n", circumference)
}

// func calculate(d float64) (float64, float64) { // parameter d adalah diameter
// 	// Hitung luas Lingkaran
// 	var area = math.Pi * math.Pow(d/2, 2)

// 	// Hitung Keliling
// 	var circumference = math.Pi * d

// 	// Mengembalikan 2 Nilai
// 	return area, circumference
// }

// Fungsi Dengan Predefined Return Value
func calculate(d float64) (area float64, circumference float64) {
	area = math.Pi * math.Pow(d/2, 2)
	circumference = math.Pi * d

	return
}
