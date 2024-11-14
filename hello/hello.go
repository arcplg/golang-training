package main

import (
	"fmt"
	"math"
)

func add(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return x, y
}

var too, ting string

func test() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	// var z uint = uint(f)
	fmt.Println(x, y, f)
}

func lap() {
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

type Vertex struct {
	X int
	Y int
}

func main() {
	// fmt.Println(add(42, 13))
	// fmt.Println(add2(42, 13))

	// a, b := swap("hello", "work")
	// fmt.Println(a, b)
	// fmt.Println(too, ting)

	// var c, teo, java = 1, 2, "toom"
	// fmt.Println(c, teo, java)

	// var i int
	// var f float64
	// var d bool
	// var s string
	// fmt.Printf("%v %v %v %q\n", i, f, d, s)

	// test()

	// lap()

	// fmt.Println(
	// 	pow(3, 2, 10),
	// 	pow(3, 3, 20),
	// )

	// defer fmt.Println("world")

	// fmt.Println("hello")

	// fmt.Println("counting")

	// for i := 0; i < 10; i++ {
	// 	defer fmt.Println(i)
	// }

	// fmt.Println("done")

	// fmt.Println(Vertex{1, 2})

	// names := [4]string{
	// 	"John",
	// 	"Paul",
	// 	"George",
	// 	"Ringo",
	// }
	// fmt.Println(names)

	// a := names[0:2]
	// b := names[1:4]

	// fmt.Println(a, b)

	// b[0] = "XXX"
	// fmt.Println(a, b)
	// fmt.Println(names)

	// s := []int{2, 3, 5, 7, 11, 13}

	// s = s[1:4]
	// fmt.Println(s)

	// s = s[0:1]
	// fmt.Println(s)

	// s = s[1:2]
	// fmt.Println(s)

	// s := []int{2, 3, 5, 7, 11, 13}
	// printSlice(s)

	// // Slice the slice to give it zero length.
	// s = s[:0]
	// printSlice(s)

	// // Extend its length.
	// s = s[:4]
	// printSlice(s)

	// // Drop its first two values.
	// s = s[2:]
	// printSlice(s)
	a := make([]int, 5)
	printSlice("a", a)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
