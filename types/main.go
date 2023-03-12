package main

import "fmt"

type ID int

var f ID = 1

func main() {
	println(f)

	fmt.Printf("O tipo de f é %T", f)  // %T Print the type
	fmt.Printf("O valor de f é %v", f) // %v Print the value
}
