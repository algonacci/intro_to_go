package main

import "fmt"

func main() {
	type NoKTP string

	var ktpEric NoKTP = "100100100"
	fmt.Println(ktpEric)
	fmt.Println(NoKTP("123123132"))
}
