package main

import (
	"fmt"
)

// tạo 1 struct để quản lý channels
type result struct {
	id  int
	sum int
}

func sum(s []int, c chan result, id int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- result{id, sum} // set sum for c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan result) // tạo 1 channel c có kiểu là result

	go sum(s[:len(s)/2], c, 1) // s[:len(s)/2] => lấy giá trị nữa đầu của array
	go sum(s[len(s)/2:], c, 2) // s[len(s)/2:] => lấy giá trị nữa sau của array
	go sum(s, c, 3)

	x, y, z := <-c, <-c, <-c // get form c

	fmt.Println(x, y, z)
}
