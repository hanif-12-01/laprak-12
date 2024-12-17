package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	if num == 0 {
		fmt.Println(0)
		return
	}

	for num > 0 {
		digit := num % 10
		fmt.Println(digit)
		num = num / 10
	}
}
