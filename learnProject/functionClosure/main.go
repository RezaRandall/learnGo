package main

import (
	"fmt"
)

// CLOSURE DISIMPAN SEBAGAI VARIABLE
func main() {
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}

	// CLOSURE MENGGUNAKAN MANIFEST TYPING
	// var getMinMaxManifestTyping (func(string, int, []string) int)
	// getMinMaxManifestTyping = func(a string, b int, c []string) {
	// 	// ..
	// }

	var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3, 1, 27}
	var min, max = getMinMax(numbers)
	fmt.Printf("data : %v\nmin : %v\nmax : %v\n", numbers, min, max)
	// Template %v digunakan untuk menampilkan segala jenis data.
	// Bisa array, int, float, bool, dan lain

	// IMMEDIATELY - INVOKED FUNCTION EXPRESSION (IIFE)
	var newNumbers = func(min int) []int {
		var r []int
		for _, e := range numbers {
			if e < min {
				continue
			}
			r = append(r, e)
		}
		return r
	}(3)
	fmt.Println("\nOriginal Number : ", numbers)
	fmt.Println("Filtered Numbers : ", newNumbers, "\n")

	//
	var maxx = 3
	var howMany, getNumbers = findmax(numbers, maxx)
	var theNumbers = getNumbers()

	fmt.Printf("Numbers : %v\t\n", numbers)
	fmt.Printf("Find numbers : %v\t\n", maxx)

	fmt.Printf("Found : %v\t\n", howMany)
	fmt.Printf("Value : %v\t\n", theNumbers)
}

func findmax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}
	return len(res), func() []int {
		return res
	}
}
