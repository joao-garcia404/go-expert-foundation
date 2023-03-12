package main

import "fmt"

func maps() {
	salaries := map[string]int{"Joao": 1000, "John": 2000}

	fmt.Println(salaries["Joao"]) // Getting a map item

	delete(salaries, "John") // Deleting a map item

	salaries["Carlos"] = 3000 // Adding a map item

	fmt.Println(salaries)

	// sal := make(map[string]int) // Creating a empty map
	// sall := map[string]int{}    // Creating a empty map

	for employee, sallary := range salaries {
		fmt.Printf("O salario do %s é de %d\n", employee, sallary)
	}
}

func slices() {
	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 100}

	fmt.Printf("len=%d, cap=%d, %v\n", len(s), cap(s), s)

	// Cleaning slice
	fmt.Printf("len=%d, cap=%d, %v\n", len(s[:0]), cap(s[:0]), s[:0])

	// Dropping slice to first 4 items
	fmt.Printf("len=%d, cap=%d, %v\n", len(s[:4]), cap(s[:4]), s[:4])

	// Getting all elements after 2 first positions
	fmt.Printf("len=%d, cap=%d, %v\n", len(s[2:]), cap(s[2:]), s[2:])

	// Adding an element
	s = append(s, 110)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func arrays() {
	var array [3]int

	array[0] = 10
	array[1] = 20
	array[2] = 30

	fmt.Println(array[len(array)-1]) // Getting array last element

	for i, v := range array {
		fmt.Printf("O valor do indice %d é %d\n", i, v)
	}
}

func main() {
	// arrays()
	slices()
	// maps()
}
