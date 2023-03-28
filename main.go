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

}
