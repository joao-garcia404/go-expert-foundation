package main

type MyNumber int

type Number interface {
	~int | float64 // With ~ MyNumber type can be accepted
}

func sum[T Number](m map[string]T) T { // Sum function with generic T
	var total T

	for _, v := range m {
		total += v
	}

	return total
}

func compare[T comparable](a T, b T) bool {
	if a == b {
		return true
	}

	return false
}

func main() {
	m1 := map[string]int{"John": 1000, "Carlos": 5000}
	m2 := map[string]float64{"John": 1000.44, "Carlos": 12.77}
	m3 := map[string]MyNumber{"John": 1000, "Carlos": 100}

	println(sum(m1))
	println(sum(m2))
	println(sum(m3))
	println(compare(10, 10))
}
