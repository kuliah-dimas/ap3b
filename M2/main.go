package main

import "fmt"

type Mahasiswa struct {
	Nama    string
	NIM     string
	Jurusan string
}

func main() {

	var jumlahMahasiswa int
	fmt.Print("Masukkan jumlah mahasiswa: ")
	fmt.Scan(&jumlahMahasiswa)

	mahasiswas := make([]Mahasiswa, jumlahMahasiswa)

	for i := 0; i < len(mahasiswas); i++ {
		fmt.Printf("Data Mahasiswa %d:\n", i+1)
		fmt.Print("Nama: ")
		fmt.Scan(&mahasiswas[i].Nama)
		fmt.Print("NIM: ")
		fmt.Scan(&mahasiswas[i].NIM)
		fmt.Print("Jurusan: ")
		fmt.Scan(&mahasiswas[i].Jurusan)
		fmt.Println()
	}

	fmt.Println("\nData Mahasiswa:")
	for i := 0; i < len(mahasiswas); i++ {
		mahasiswaPtr := &mahasiswas[i]
		fmt.Printf("Nama: %s\n", mahasiswaPtr.Nama)
		fmt.Printf("NIM: %s\n", mahasiswaPtr.NIM)
		fmt.Printf("Jurusan: %s\n", mahasiswaPtr.Jurusan)
		fmt.Println("-")
	}
}
