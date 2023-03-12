package main

import "fmt"

type Adress struct {
	street string
	number int
	city   string
	state  string
}

type Client struct {
	name   string
	age    int
	active bool
	adress Adress
}

type Person interface {
	disableClient()
}

func (c Client) disableClient() {
	c.active = false

	fmt.Printf("O cliente %s foi desativado\n", c.name)
}

func disable(person Person) {
	person.disableClient()
}

func main() {
	client := Client{
		name:   "John",
		age:    20,
		active: true,
	}

	// client.active = false
	// client.disableClient() // Disabling by struct method

	disable(client)

	fmt.Printf(
		"Nome: %s, Idade: %d, Ativo: %t", client.name, client.age, client.active,
	)
}
