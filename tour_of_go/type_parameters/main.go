package main

import "fmt"

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}

	return -1
}

func main() {
	si := []int{15, 20, 25, 30}
	fmt.Println(Index(si, 20))

	ss := []string{"Thien", "PHP", "MYSQL", "NODE"}
	fmt.Println(Index(ss, "NODEJS"))
}
