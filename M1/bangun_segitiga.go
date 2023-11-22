package main

import "fmt"

func main() {
	var tinggi int
	fmt.Print("Masukkan tinggi: ")
	fmt.Scan(&tinggi)

	fmt.Println("Pola segitiga")
	for i := 1; i <= tinggi; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println("")
	}

	fmt.Println("Pola Berbalik")
	for i := tinggi; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println("")
	}
}
