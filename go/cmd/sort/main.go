package main

import (
	"fmt"

	"example.com/go-algo/pkg/quicksort"
)

func main() {
	a := []int{2, 13, 2, 5, 612, 22, 12}
	quicksort.Sort(a)
	fmt.Printf("a: %v\n", a)
}
