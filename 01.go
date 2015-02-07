package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	var c int
	var na int
	var nb int
	var nab int
	var sa int
	var sb int
	var sab int
	var sum int
	
	a = 3
	b = 5
	c = 1000000000

	//mencari jumlah deret kelipatan 3:
	na = (c - 1) / a
	nb = (c - 1) / b
	nab = c / (a * b)

	sa = (na * (a + (a * na))) / 2
	sb = (nb * (b + (b * nb))) / 2
	sab = (nab * ((a * b) + (a * b * nab))) / 2

	sum = sa + sb - sab

	fmt.Printf("Jumlah = %d", sum)
}
