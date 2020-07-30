package main

import (
	"fmt"
	"strings"
)

// DEKLARASI STRUCT
type student struct {
	name  string
	grade int
}

// PENERAPAN STRUCT
func main() {
	var s1 student
	s1.name = "jhono koko"
	s1.grade = 2

	fmt.Println("name : \t", s1.name)
	fmt.Println("grade : ", s1.grade)

	// INISIALISASI OBJECT STRUCT
	var strata1 = student{}
	strata1.name = "wick"
	strata1.grade = 3

	var strata2 = student{"reza", 4}
	var strata3 = student{grade: 10}
	/*
		Keistimewaan lain menggunakan cara ini adalah penentuan
		nilai property bisa dilakukan dengan tidak berurutan.
	*/
	var strata4 = student{grade: 5, name: "afi"}
	var strata5 = student{name: "joule", grade: 6}

	fmt.Println("student 1 name:", strata1.name, "grade:", strata1.grade)
	fmt.Println("student 2 name:", strata2.name, "grade:", strata2.grade)
	fmt.Println("student 3 name:", strata3.name, "grade:", strata3.grade)
	fmt.Println("student 4 name:", strata4.name, "grade:", strata4.grade)
	fmt.Println("student 5 name:", strata5.name, "grade:", strata5.grade)
	fmt.Println()

	// VARIABLE OBJECT POINTER
	var s6 = student{name: "lala", grade: 7}
	var s7 *student = &s6
	fmt.Println("student 6 name :", s6.name)
	fmt.Println("student 7 name :", s7.name)

	s7.name = "ethan"
	fmt.Println("student 6 name :", s6.name)
	fmt.Println("student 6 name :", s6.name)
	fmt.Println("student 7 name :", s7.name)
	fmt.Println("student 6 name :", s7.name)

	// EMBEDED STRUCT
	var sd = pelajar{}
	sd.name = "bejo"
	sd.age = 45
	sd.grade = 3

	fmt.Println("name", sd.name)
	fmt.Println("age", sd.age)
	fmt.Println("age", sd.orang.age)
	fmt.Println("grade", sd.grade)

	// EMBEDED STRUCT DENGAN NILAI PROPERTY YANG SAMA
	var ss = murid{}
	ss.name = "Lika"
	ss.manusia.age = 20 // age dari struct manusia
	ss.age = 10         // age dari struct pelajar

	fmt.Println()
	fmt.Println("name : \t", ss.name)
	fmt.Println("age dari struct orang  : \t", ss.manusia.age)
	fmt.Println("age dari struct pelajar  : \t", ss.age)

	// PENGISIAN NILAI SUB - STRUCT
	var c1 = car{nama: "ferrari", code: 01}
	var b1 = bahanBakar{car: c1, jenisBBM: "Pertamax"}

	fmt.Println()
	fmt.Println("name : \t\t", b1.nama)
	fmt.Println("code : \t\t", b1.code)
	fmt.Println("jenis bbm : \t", b1.jenisBBM)

	// ANONYMOUS STRUCT
	var sc1 = struct {
		superCar
		gradeCar int
	}{}
	sc1.carName = "Lamborghini"
	sc1.codeCar = 666
	sc1.gradeCar = 13

	fmt.Println()
	fmt.Println("Name : \t\t", sc1.superCar.carName)
	fmt.Println("Car Code : \t", sc1.superCar.codeCar)
	fmt.Println("Grade Car : \t", sc1.gradeCar)

	// KOMBINASI SLICE & STRUCT
	var mt = []motor{
		{motorType: "Supra", motorCode: 111},
		{motorType: "Jupiter", motorCode: 222},
		{motorType: "Beat", motorCode: 333},
	}

	fmt.Println()
	for _, m := range mt {
		fmt.Println(m.motorType, "The code is : ", m.motorCode)
	}

	// INISIALISASI SLICE ANONYMOUS STRUCT
	var sc2 = []struct {
		superCar
		madeIn string
	}{
		{superCar: superCar{"audi", 99}, madeIn: "london"},
		{superCar: superCar{"toyota", 88}, madeIn: "japan"},
		{superCar: superCar{"ford", 77}, madeIn: "USA"},
	}
	fmt.Println()
	for i, supCar := range sc2 {
		i++
		fmt.Println(i, supCar.carName, supCar.codeCar, supCar.madeIn)
	}

	// DEKLARASI ANONYMOUS STRUCT MENGGUNAKAN KEYWORD var
	motorCycle.motor = motor{motorType: "bemo", motorCode: 928139}
	motorCycle.engine = "V8"
	fmt.Println()
	fmt.Println("Motor type :", motorCycle.motorType, ", this motor code is : ", motorCycle.motorCode, ", and the engin :", motorCycle.engine)

	fmt.Println("")
	// HANYA DEKLARASI (yang artinya dicetak sebuah objek dari anonymous struct kemudian disimpan pada variabel bernama motorCycle)
	footballPlayer.playerName = "Diego"
	fmt.Println(footballPlayer.playerName, baseBallPlayer.playerNo)

	fmt.Println()
	// NESTED DTRUCT
	var fb = football{grade: 44, hobbies: []string{"Hikking", "cycling", "running"}}
	// fb.person.name = "Boby"
	// fb.person.age = 26
	printHobbies(fb.hobbies)

	fmt.Println()
	// TYPE ALIAS
	type ban = tires
	var t1 = tires{"dunlop", 17}
	fmt.Println(t1)

	var t2 = ban{"maxxis", 89}
	fmt.Println(t2)

	Ban := ban{"gt", 45}
	fmt.Println(tires(Ban))

	Tires := tires{"gt", 45}
	fmt.Println(ban(Tires))

}
func printHobbies(arrHobbies []string) {
	var fb = football{grade: 44, hobbies: []string{"Hikking", "cycling", "running"}}
	fb.person.name = "Boby"
	fb.person.age = 26

	var hobbiesName = strings.Join(arrHobbies, ", ")
	fmt.Print("football player :", fb.person.name, ", age ", fb.person.age, ", and his hobby are ", hobbiesName)
}

// EMBEDED STRUCT
type orang struct {
	name string
	age  int
}

type pelajar struct {
	orang
	grade int
}

// EMBEDED STRUCT DENGAN NILAI PROPERTY YANG SAMA
type manusia struct {
	name string
	age  int
}

type murid struct {
	manusia
	grade int
	age   int
}

// PENGISIAN NILAI SUB-STUCT
type car struct {
	nama string
	code int
}

type bahanBakar struct {
	jenisBBM string
	car
}

// ANONYMOUS STRUCT
type superCar struct {
	carName string
	codeCar int
}

// KOMBINASI SLICE DAN STRUCT
type motor struct {
	motorType string
	motorCode int
}

// DEKLARASI ANONYMOUS STRUCT MENGGUNAKAN KEYWORD var
var motorCycle struct { // HANYA DEKLARASI (yang artinya dicetak sebuah objek dari anonymous struct kemudian disimpan pada variabel bernama motorCycle)
	motor
	engine string
}

var footballPlayer struct { // HANYA DEKLARASI (yang artinya dicetak sebuah objek dari anonymous struct kemudian disimpan pada variabel bernama motorCycle)
	playerName string
}

var baseBallPlayer = struct { // DEKLARASI SEKALIGUS INISIALISASI
	playerNo int
}{
	99,
}

// NEXT NESTED STRUCT
type football struct {
	person struct {
		name string
		age  int
	}
	grade   int
	hobbies []string
}

// DEKLARASI DAN INISIALISASI STRUCT SECARA HORIZONTAL
type fruits struct {
	name  string
	color string
	code  int
} // semi-colon (;) digunakan sebagai pembatas deklarasi poperty yang dituliskan secara horizontal
var fish = struct {
	name  string
	color string
	code  int
}{name: "nemo", color: "red", code: 23}
var animal = struct {
	name  string
	types string
	code  int
}{"buaya", "darat", 11}

// TAG PROPERTY DALAM STRUCT
type bearing struct {
	name     string `tag1`
	diameter int    `tag2`
}

// TYPE ALIAS
type tires struct {
	name     string
	diameter int
}

type home struct {
	name  string
	large int
}
