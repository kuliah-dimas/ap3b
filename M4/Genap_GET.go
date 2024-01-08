package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Mahasiswa
type Mahasiswa struct {
	ID      int    `json:"id"`
	NIM     int    `json:"nim"`
	Name    string `json:"name"`
	Kelas   string `json:"kelas"`
	Jurusan string `json:"jurusan"`
}

// NewMahasiswa
func NewMahasiswa() []Mahasiswa {
	mhs := []Mahasiswa{
		Mahasiswa{
			ID:      1,
			NIM:     50422430,
			Name:    "Dimas Febriyanto",
			Kelas:   "2IA21",
			Jurusan: "IT",
		},
		Mahasiswa{
			ID:      2,
			NIM:     724,
			Name:    "can i be ur bf?",
			Kelas:   "2IA21",
			Jurusan: "IT",
		},
	}
	return mhs
}

// GetMahasiswa
func GetMahasiswa(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		mhs := NewMahasiswa()
		datamahasiswa, err := json.Marshal(mhs)
		if err != nil {
			http.Error(w, err.Error(),
				http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(datamahasiswa)
		return
	}

	http.Error(w, "hayo mau ngapain", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/mahasiswa", GetMahasiswa)
	fmt.Println("server running...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
