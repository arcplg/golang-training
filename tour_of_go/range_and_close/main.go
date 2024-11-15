package main

import (
	"fmt"
	"range_and_close/fibonacci"
	"range_and_close/testModule"
)

func main() {
	fmt.Println("check OK")

	c := make(chan int, 10)
	go fibonacci.Generate(5, c)
	for i := range c {
		fmt.Println(i)
	}

	c2 := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci.FibonacciSelectCheck(c2, quit)

	testModule.ShowTextOk()
}
