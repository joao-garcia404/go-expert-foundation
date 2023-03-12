package main

func sum(a, b *int) int {
	*a = 50
	*b = 50

	return *a + *b
}

func main() {
	a := 10

	var pointer *int = &a
	*pointer = 20

	b := &a // Assigning a memory adress to b

	*b = 30 // Changes the value of a as well

	var1 := 10
	var2 := 20

	sum(&var1, &var2)
	println(var1)
	println(var2)
}
