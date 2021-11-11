package main

import "fmt"

func main() {
	var name1 = "Eric"
	var name2 = "Julianto"

	var result bool = name1 != name2

	fmt.Println(result)

	var value1 = 100
	var value2 = 200

	fmt.Println(value1 > value2)  // FALSE
	fmt.Println(value1 < value2)  // TRUE
	fmt.Println(value1 == value2) // FALSE
	fmt.Println(value1 != value2) // TRUE
}
