package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	
	var (
		umur int
		alamat string
		isMarried bool
	)

	umur = 10
	alamat = "tambun bekasi"
	isMarried = false

	fmt.Println("umur", umur, "alamat", alamat, "udah nikah?", isMarried)

	// variabel harus dipakai, jika tidak maka bisa di ignore dengan tanda underscore
	test1, _ := 10, 10
	fmt.Println(test1)

	// PascalCase atau camelCase variable name
	MyFamily := 10
	myAddress := "teks"
	fmt.Println(myAddress)
	fmt.Println(MyFamily)

	// convert data type
	var angka int32 = 8
	
	var bigNum = int64(angka)
	var floatNum = float64(bigNum)
	fmt.Println(floatNum)

	// convert angka ke string
	var hurufA int8 = 65
	s := fmt.Sprint(hurufA)
	// b := string(hurufA)
	fmt.Println(s)

	// condition simple
	score := 80
	if score > 75 {
		fmt.Println("Good")
	}

	// condition if else 
	if score > 75 {
		fmt.Println("Good")
	} else {
		fmt.Println("Bad")
	}
	
	// condition if else if
	if score > 75 {
		fmt.Println("Good")
	} else if score < 75 {
		fmt.Println("Bad")
	} else {
		fmt.Println("So Bad wkwk")
	}
	fmt.Println("End")

	// condition bersarang atau nested condition
	menikah := false
	umurnya := 25

	if umurnya > 25 {
		if menikah {
			fmt.Println("You are married and over 25")
		} else {
			fmt.Println("You are not married and over 25")
		}
	} else {
		fmt.Println("You are not over 25")
	}
}
