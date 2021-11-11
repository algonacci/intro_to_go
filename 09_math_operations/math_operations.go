package main

import "fmt"

func main() {
	var a = 10
	var b = 5
	var c = a * b
	fmt.Println(c)

	var result = 10 + 5
	fmt.Println(result)

	var i = 100
	i += 100 // Sama saja seperti i = i + 100

	fmt.Println(i)

	i++ // Sama saja seperti i = i + 1
	fmt.Println(i)
}
