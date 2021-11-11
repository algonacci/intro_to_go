package main

import "fmt"

func main() {
	var name string

	name = "Eric Julianto"
	fmt.Println("Nama: ", name)

	var pekerjaan = "Developer"
	var umur = 20

	fmt.Println("Pekerjaan: ", pekerjaan)
	fmt.Println("Umur: ", umur)

	semester := 5
	country := "Indonesia"

	fmt.Println("Semester kuliah: ", semester)
	fmt.Println("Tinggal di: ", country)

	var (
		jurusan1 = "Hospitality"
		jurusan2 = "Pariwisata"
	)

	fmt.Println("Jurusan Kuliah: ",jurusan1," dan ", jurusan2)

	first,mid,last := "Universitas ", "Bunda ", "Mulia"
	fmt.Println("Kuliah:", first, mid, last)
}