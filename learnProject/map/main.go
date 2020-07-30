package main

import "fmt"

func main() {
	var chicken map[string]int
	chicken = map[string]int{}

	chicken["January"] = 50
	chicken["February"] = 100

	fmt.Println("January", chicken["January"])
	fmt.Println("December", chicken["December"])
	fmt.Println("February", chicken["February"], "\n")

	var cat map[int]string

	cat = map[int]string{}

	cat[3] = "Mia"
	cat[1] = "ciko"
	cat[2] = "bubul"

	fmt.Println("cat 1", cat[1])
	fmt.Println("cat 2", cat[2])

	dog := map[int]string{}
	dog[1] = "blackie"
	fmt.Println("This Dog name is ", dog[1])

	// CARA VERTIKAL
	var dogOne = map[int]string{1: "Cihua-hua", 2: "Kintamani"}
	dogTwo := map[int]string{1: "Kintamani", 2: "Scooby"}

	fmt.Println("Dog One", dogOne[1], dogOne[2])
	fmt.Println("Dog Two", dogTwo[1], dogTwo[2])

	// CARA HORIZONTAL
	var dogThree = map[int]string{
		1: "Golden",
		2: "Pitbul",
	}
	fmt.Println("Dog Three", dogThree[1], dogThree[2])

	dogFour := map[int]string{
		1: "Kikik",
		2: "Kribo",
	}
	fmt.Println("Dog Four", dogFour[1], dogFour[2])

	// ITERASI ITEM MAP MENGGUNKAN for - range
	for key, val := range dogOne {
		fmt.Println(key, " \t", val)
	}

	// MENGHAPUS ITEM map
	fmt.Println("")
	fmt.Println(len(dogOne))
	fmt.Println(dogOne)

	delete(dogOne, 1)
	fmt.Println(len(dogOne))
	fmt.Println(dogOne, "\n")

	// DETEKSI KEBERADAAN ITEM DENGAN KEY TERTENTU
	var value, isExsist = dogFour[3]

	if isExsist {
		fmt.Println(value)
	} else {
		fmt.Println("items did not exsist")
	}

	// KOMBINASI slice & map
	fishs := []map[string]string{
		map[string]string{"name": "Gatul", "type": "Herbivora", "habitat": "Air Tawar"},
		map[string]string{"name": "Mujaer", "type": "Omnivora", "habitat": "Air Tawar"},
		map[string]string{"name": "Hiu", "type": "Karnivora", "habitat": "Air Laut"},
	}
	for _, fish := range fishs {
		fmt.Println(fish["name"], fish["type"], fish["habitat"])
	}
	// DIATAS ADALAH PENJABARAN slice & map VERSI LAMA
	// JIKA MENGGUNAKAN GO VERSI TERBARU BISA DIPERSINGKAT SEPERTI DIBAWAH INI

	ikans := []map[string]string{
		{"name": "Gatul", "type": "Herbivora", "habitat": "Air Tawar"},
		{"name": "Mujaer", "type": "Omnivora", "habitat": "Air Tawar"},
		{"name": "Hiu", "type": "Karnivora", "habitat": "Air Laut"},
	}
	for _, ikan := range ikans {
		fmt.Println(ikan["name"], ikan["type"], ikan["habitat"])
	}

	// Dalam []map[string]string, tiap elemen bisa saja memiliki key yang berbeda-beda, sebagai contoh seperti kode berikut.
	var data = []map[string]string{
		{"name": "chicken blue", "gender": "male", "color": "brown"},
		{"address": "mangga street", "id": "k001"},
		{"community": "chicken lovers"},
	}

	for _, dataData := range data {
		fmt.Println(dataData["name"], dataData["gender"], dataData["color"], dataData["name"], dataData["address"], dataData["id"], dataData["community"])
	}

}
