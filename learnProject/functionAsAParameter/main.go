package main

import (
	"fmt"
	"strings"
)

// FUNGSI SEBAGAI PARAMETER
func main() {
	var data = []string{"wick", "jason", "ethan", "mamen", "rachel", "betty"}
	var dataConstainsO = filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})
	var dataLength5 = filter(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println("Data asli \t\t:", data)
	fmt.Println("Filter ada huruf \"o\"\t:", dataConstainsO)
	fmt.Println("Filter jumlah huruf \"5\"\t:", dataLength5)

	// HASIL FILTER MENGGUNKAN ALIAS
	dataConstainY := filterTwo(data, func(eachData string) bool {
		return strings.Contains(eachData, "y")
	})
	dataLenght6 := filterTwo(data, func(eachData string) bool {
		return len(eachData) == 6
	})

	fmt.Println()
	fmt.Println("Original Data \t\t:", data)
	fmt.Println("Filter ada huruf \"y\"\t:", dataConstainY)
	fmt.Println("Filter jumlah huruf \"6\"\t:", dataLenght6)

	/*
		Penggunaan Fungsi string.Contains()
		Inti dari fungsi ini adalah untuk deteksi apakah sebuah substring adalah bagian dari string,
		jika iya maka akan bernilai true, dan sebaliknya. Contoh penggunaannya seperti dibawah ini
		Variabel result bernilai true karena string "ang" merupakan bagian dari string "Golang"
	*/
	var results = strings.Contains("Betty", "y") // penggunaan Contains menghasilkan nilai boolean
	fmt.Println()
	fmt.Println(results)
}

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}

// FILTER KE 2 MENGGUNAKAN ALIAS
type functionCallback func(string) bool

func filterTwo(data []string, callback functionCallback) []string {
	var result []string
	for _, eachData := range data {
		if filtered := callback(eachData); filtered {
			result = append(result, eachData)
		}
	}
	return result
}
