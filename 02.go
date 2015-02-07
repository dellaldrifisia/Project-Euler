package main

import (
	"fmt"
)

func main() {
	var x [100]int
	var total int
	total = 0

	//coba tanpa nilai 35
	loop := true
	val1 := 0

	for loop {
		val1++

		if x[val1]%2 == 0 && x[val1] <= 4000000 {
			total += x[val1]
			loop = true
		}

		if x[val1] > 4000000 {
			loop = false
		}
	}

	fmt.Println(x)
	fmt.Printf("%d", total)

}
