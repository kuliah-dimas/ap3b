package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Mahasiswa
type Biodata struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Umur   int    `json:"umur"`
	Alamat string `json:"alamat"`
	Nohp   string `json:"no_hp"`
}

// NewMahasiswa
func NewBiodata() []Biodata {
	mhs := []Biodata{
		Biodata{
			ID:     1,
			Name:   "Dimas Febriyanto",
			Umur:   19,
			Alamat: "Jl. Bersamamu, aku nyaman",
			Nohp:   "082329135125",
		},
		Biodata{
			ID:     2,
			Name:   "my gf soon",
			Umur:   19,
			Alamat: "Sama, Jl. with u Iam happy",
			Nohp:   "08 berapa?",
		},
	}
	return mhs
}

// GetMahasiswa
func GetBiodatadiri(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		mhs := NewBiodata()
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
	http.HandleFunc("/biodatadiri", GetBiodatadiri)
	fmt.Println("server running...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
