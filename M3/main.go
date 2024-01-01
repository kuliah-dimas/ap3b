package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/mahasiswa")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database")

	var nama, umur, alamat string
	fmt.Print("Masukkan Nama: ")
	fmt.Scanln(&nama)
	fmt.Print("Masukkan Umur: ")
	fmt.Scanln(&umur)
	fmt.Print("Masukkan Alamat: ")
	fmt.Scanln(&alamat)

	_, err = db.Exec("INSERT INTO biodata (nama, umur, alamat) VALUES (?, ?, ?)", nama, umur, alamat)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Biodata inserted successfully")

	rows, err := db.Query("SELECT nama, umur, alamat FROM biodata")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Biodata from the database:")
	for rows.Next() {
		var nama, umur, alamat string
		err := rows.Scan(&nama, &umur, &alamat)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Nama: %s, Umur: %s, Alamat: %s\n", nama, umur, alamat)
	}
}
