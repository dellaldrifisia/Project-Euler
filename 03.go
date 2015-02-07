package main

import (
	"fmt"
)

func main() {
	var i int
	j := 0
	var x int
	x = 20
	a := make([]int, x+1)

	for i := 1; i <= x; i++ {
		if x%i == 0 {
			j++
		}
	}

	if j == 2 {
		a[i] = i
	}

	fmt.Println(a)
}
