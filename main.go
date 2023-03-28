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
}
