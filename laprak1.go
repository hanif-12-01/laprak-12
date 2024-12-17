package main

import "fmt"

func main() {
	var username, password string
	var gagal int

	for {
		fmt.Println("masukkan username dan pasword:")
		fmt.Scan(&username, &password)

		if username == "Admin" && password == "Admin" {
			fmt.Printf("%d percobaan gagal login\n", gagal)
			break
		}
		gagal++
	}
}
