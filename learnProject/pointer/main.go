package main

import "fmt"

// PENERAPAN POINTER
func main() {
	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (value)\t\t", numberA)
	fmt.Println("numberA (address)\t", &numberA)
	fmt.Println("numberB (value)\t\t", *numberB)
	fmt.Println("numberB (address)\t", numberB)

	// EFEK PERUBAHAN NILAI POINTER
	fmt.Println()

	numberA = 5

	fmt.Println("numberA (value)\t\t", numberA)
	fmt.Println("numberA (address)\t", &numberA)
	fmt.Println("numberB (value)\t\t", *numberB)
	fmt.Println("numberB (address)\t", numberB)

	// PARAMETER POINTER
	var numA = 4
	fmt.Println("Before : ", numA)
	fmt.Println("Before (Address): ", &numA)
	fmt.Println()

	change(&numA, 10)
	fmt.Println("After :", numA)
	fmt.Println("After (Address): ", &numA)
}

func change(original *int, valAfter int) {
	*original = valAfter
}
