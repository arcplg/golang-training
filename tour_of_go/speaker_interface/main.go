package main

import (
	"fmt"
)

type Speaker interface {
	Speaker() string
	Introduce() string
}

type Person struct {
	Name string
}

func (p Person) Speaker() string {
	return "Hello!"
}

func (p Person) Introduce() string {
	return "My name is " + p.Name
}

type Dog struct{}

func (d Dog) Speaker() string {
	return "Woof!"
}

func (d Dog) Introduce() string {
	return "I'm a dog."
}

func saySomeThing(s Speaker) {
	fmt.Println(s.Speaker())
	fmt.Println(s.Introduce())
}

func main() {
	p := Person{Name: "Thien"}
	d := Dog{}

	saySomeThing(p)
	saySomeThing(d)
}
