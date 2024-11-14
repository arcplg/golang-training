package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)\n", p.Name, p.Age)
}

func main() {
	a := Person{"Thien 1", 27}
	b := Person{"Thien 2", 27}
	fmt.Println(a, b)
	fmt.Println("check ok")
}
