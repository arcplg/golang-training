package main

import "fmt"

func main() {
	checkString()
	checkInt()
}

func checkString() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// case error
	// f = i.(float64)
	// fmt.Println(f)

	fmt.Println("OK")
}

func checkInt() {
	var i interface{} = 1

	s := i.(int)
	println(s)

	s, ok := i.(int)
	println(s, ok)

	f, ok := i.(string)
	println(f, ok)
}
