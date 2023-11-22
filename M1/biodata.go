package main

import "fmt"

func main() {
	var (
		nama    string
		kelas   string
		npm     string
		jurusan string
	)

	fmt.Print("Masukkan nama: ")
	fmt.Scan(&nama)
	fmt.Print("Masukkan kelas: ")
	fmt.Scan(&kelas)
	fmt.Print("Masukkan npm: ")
	fmt.Scan(&npm)
	fmt.Print("Masukkan jurusan: ")
	fmt.Scan(&jurusan)

	fmt.Println("Halo nama saya", nama, "saya dari kelas", kelas, "dengan npm", npm, "saya dari jurusan", jurusan)
}
