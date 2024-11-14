package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
	"strings"
	"time"
)

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	return
}

var c, python, java bool

// Variables with initializers
var j, k int = 1, 2

// Basic types
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// START Numeric Constants
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

// END Numeric Constants

// IF
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}

	return fmt.Sprintln(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}

	return lim
}

// Exercise: Loops and Functions

// Structs
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Printf("Hello, Thien")

	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())

	fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Println("Now you have %g problems .\n", math.Sqrt(7))

	fmt.Println(math.Pi)

	fmt.Println(add(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(20))

	var i int
	fmt.Println(i, c, python, java)

	// Variables with initializers
	var c, python, java = true, false, "no!"
	fmt.Println(j, k, c, python, java)

	// Basic types
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe) // T = type, v = value
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// Type conversions
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	// Type inference
	v := 42 // change me!
	fmt.Printf("v is of type %T\n", v)

	// Numeric Constants
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	// For
	sum := 0
	for o := 0; o < 10; o++ {
		sum += o
	}
	fmt.Println(sum)

	sumTemp := 0
	for q := 0; q < 5; q++ {
		sumTemp += q
	}
	fmt.Println("sum temp:", sumTemp)

	// For continued
	sum = 50
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// IF
	fmt.Println(sqrt(2), sqrt(-4))

	// If with a short statement
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	// SWITCH
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("----------------- Switch evaluation order -----------------")

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	fmt.Println("----------------- Switch with no condition -----------------")

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	fmt.Println("----------------- Defer -----------------")
	defer fmt.Println("world")

	fmt.Println("hello")

	fmt.Println("----------------- Stacking defers -----------------")
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")

	fmt.Println("----------------- Pointers -----------------")
	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)

	fmt.Println("----------------- Structs -----------------")
	fmt.Println(Vertex{1, 2})

	fmt.Println("----------------- Struct Fields -----------------")
	vv := Vertex{1, 2}
	fmt.Println(vv.X)
	vv.X = 4
	fmt.Println(vv.X)

	fmt.Println("----------------- Pointers to structs -----------------")
	vvv := Vertex{1, 2}
	ppp := &vvv
	ppp.X = 1e9
	fmt.Println(vvv)

	fmt.Println("----------------- Struct Literals -----------------")
	fmt.Println(v1, p1, v2, v3)

	fmt.Println("----------------- Arrays -----------------")
	var a1 [2]string
	a1[0] = "Hello"
	a1[1] = "World"
	fmt.Println(a1[0], a1[1])
	fmt.Println(a1)

	fmt.Println("----------------- Slices -----------------")
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	fmt.Println("----------------- Slices are like references to arrays -----------------")
	names1 := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names1)
	a2 := names1[1:2]
	b1 := names1[1:3]
	fmt.Println(a2, b1)

	b1[0] = "XXX"
	fmt.Println(a1, b1)
	fmt.Println(names1)

	fmt.Println("----------------- Slice literals -----------------")
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	fmt.Println("----------------- Slice defaults -----------------")
	s1 := []int{2, 3, 5, 7, 11, 13}
	s1 = s1[1:4]
	fmt.Println(s1)

	s1 = s1[:4]
	fmt.Println(s1)

	s1 = s1[1:]
	fmt.Println(s1)

	fmt.Println("----------------- Slice length and capacity -----------------")
	sliceLengthAndCapacity()

	fmt.Println("----------------- Creating a slice with make -----------------")
	createSliceWithMake()

	fmt.Println("----------------- Slices of slices	-----------------")
	sliceOfSlice()

	fmt.Println("----------------- Appending to a slice -----------------")
	appendingToSlice()

	fmt.Println("----------------- Range -----------------")
	rangePow()

	fmt.Println("----------------- Range continued -----------------")
	rangeContinued()

	fmt.Println("----------------- Maps -----------------")
	mapsTour()

	fmt.Println("----------------- Map Literals -----------------")
	mapLiterals()

	fmt.Println("----------------- Mutating Maps -----------------")
	mutatingMaps()

	fmt.Println("----------------- Function closures -----------------")
	fucntionClouseres()

	fmt.Println("----------------- Method tour -----------------")
	mothodTour()
	fmt.Println("----------------- End method tour -----------------")

	fmt.Println("----------------- Methods continued -----------------")
	fx := MyFloat(-math.Sqrt2)
	fmt.Println(fx.Abs())
	fmt.Println("----------------- Methods continued -----------------")

	fmt.Println("----------------- Pointer receivers -----------------")
	// fx := MyFloat(-math.Sqrt2)
	// fmt.Println(fx.Abs())
	fmt.Println("----------------- Pointer receivers -----------------")
	mkdir
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p1 = &Vertex{1, 2}
)

// ----------------- Slice length and capacity -----------------
func sliceLengthAndCapacity() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// ----------------- Creating a slice with make -----------------

func printCreateSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

// Creating a slice with make
func createSliceWithMake() {
	a := make([]int, 5)
	printCreateSlice("a", a)

	b := make([]int, 0, 5)
	printCreateSlice("b", b)

	c := b[:2]
	printCreateSlice("c", c)

	d := c[2:5]
	printCreateSlice("d", d)
}

// Slices of slices
func sliceOfSlice() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

// Appending to a slice
func appendingToSlice() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

// Range
func rangePow() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("i = %d \n", i)

		fmt.Printf("v = %d \n", v)

		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func rangeContinued() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}

	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

type VertexMaps struct {
	Lat, Long float64
}

// maps
func mapsTour() {
	var m map[string]VertexMaps
	m = make(map[string]VertexMaps)
	m["Bell Labs"] = VertexMaps{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

// map literals
func mapLiterals() {
	var m = map[string]VertexMaps{
		"Bell Labs": VertexMaps{
			40.68433, -74.39967,
		},
		"Google": VertexMaps{
			37.42202, -122.08408,
		},
	}

	fmt.Println(m)
}

// Mutating Maps
func mutatingMaps() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

// Function closures
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fucntionClouseres() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

// Methods
type VertexMethod struct {
	X, Y float64
}

func (v VertexMethod) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func mothodTour() {
	v := VertexMethod{3, 4}
	fmt.Println(v.Abs())
}

// methods continued
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}
