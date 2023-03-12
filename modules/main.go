package main

import (
	"fmt"
	"go-expert/maths"

	"github.com/google/uuid"
)

func main() {
	total := maths.SumValues(10, 20)

	fmt.Printf("Resultado: %v\n", total)
	fmt.Println(uuid.New())
}
