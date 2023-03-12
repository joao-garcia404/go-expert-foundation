package main

import (
	"errors"
	"fmt"
)

func main() {
	value, err := sum(1, 10)

	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(value)

	totalSum := sumNumbers(10, 20, 30, 40, 50, 60)

	fmt.Println(totalSum)

	// Closures
	total := func() int {
		return sumNumbers(10, 20, 30, 40, 50, 60) * 2
	}()

	fmt.Println(total)
}

func sum(a, b int) (int, error) {
	if a+b >= 50 {
		return 0, errors.New("A soma é maior que 50")
	}

	return a + b, nil
}

func sumNumbers(numbers ...int) int { // Getting inderteminate function args
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}
