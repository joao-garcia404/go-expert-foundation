package main

func main() {
	for i := 0; i < 10; i++ { // Basic for
		// println(i)
	}

	numbers := []string{"1", "2", "3"}

	for key, value := range numbers { // Transversing numbers slice
		println(key, value)
	}

	i := 0

	for i < 10 { // For increment
		// println(i)
		i++
	}

	for { // infinite loop
		// Do something
	}
}
