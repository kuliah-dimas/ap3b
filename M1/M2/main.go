package main

import "fmt"

type Alamat struct {
	Jalan string
	Kota  string
	Prov  string
	Pos   string
}

type Mahasiswa struct {
	Nama    string
	NIM     string
	Jurusan string
	Alamat  Alamat
}

func main() {
	var mhs Mahasiswa
	fmt.Print("Masukkan Nama Mahasiswa: ")
	fmt.Scan(&mhs.Nama)
	fmt.Print("Masukkan NIM Mahasiswa: ")
	fmt.Scan(&mhs.NIM)
	fmt.Print("Masukkan Jurusan Mahasiswa: ")
	fmt.Scan(&mhs.Jurusan)
	fmt.Print("Masukkan Alamat Jalan Mahasiswa: ")
	fmt.Scan(&mhs.Alamat.Jalan)
	fmt.Print("Masukkan Alamat Kota Mahasiswa: ")
	fmt.Scan(&mhs.Alamat.Kota)
	fmt.Print("Masukkan Alamat Provinsi Mahasiswa: ")
	fmt.Scan(&mhs.Alamat.Prov)
	fmt.Print("Masukkan Alamat Kode Pos Mahasiswa: ")
	fmt.Scan(&mhs.Alamat.Pos)

	fmt.Printf("Nama: %s\n", mhs.Nama)
	fmt.Printf("NIM: %s\n", mhs.NIM)
	fmt.Printf("Jurusan: %s\n", mhs.Jurusan)
	fmt.Printf("Alamat: \n Jalan: %s\n Kota: %s\n Provinsi: %s\n Kode Pos: %s\n", mhs.Alamat.Jalan, mhs.Alamat.Kota, mhs.Alamat.Prov, mhs.Alamat.Pos)
}
