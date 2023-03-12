package main

func main() {
	var name interface{} = "John"

	println(name)
	println(name.(string))

	res, ok := name.(int)

	println(res)
	println(ok)
}
