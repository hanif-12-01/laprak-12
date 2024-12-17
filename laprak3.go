package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	hasil := 0
	for x >= y {
		x -= y
		hasil++
	}
	fmt.Print(hasil)
}
