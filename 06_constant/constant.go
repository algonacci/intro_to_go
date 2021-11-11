package main

import "fmt"

func main() {
	const firstName string = "Eric"
	const lastName = "Julianto"

	const (
		umur      = 20
		pekerjaan = "Developer"
	)

	// PASTI ERROR KARENA CONSTANT
	// firstName = "Tidak Bisa Diubah"
	// lastName = "Tidak Bisa Diubah"

	// CONSTANT tidak kenapa-napa kalau tidak dipakai

	fmt.Println(firstName + " " + lastName)
	fmt.Println(umur)
	fmt.Println(pekerjaan)
}
