package main

import "fmt"

func sum(array []int) int {
	s := 0
	for _, item := range array {
		s += item
	}
	return s
}

func main() {
	array := []int{1, 2, 3, 412, 312, 4, 123, 123, 1412, 3, 123, 213412, 312, 312, 312, 312, 3, 124, 12, 3, 12}
	fmt.Println("Sum: ", sum(array))
}
