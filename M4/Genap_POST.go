package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Mahasiswa
type Mahasiswa struct {
	ID      int    `json:"id"`
	NIM     int    `json:"nim"`
	Name    string `json:"name"`
	Kelas   string `json:"kelas"`
	Jurusan string `json:"jurusan"`
}

// PostMahasiswa
func PostMahasiswa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Mhs Mahasiswa

	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&Mhs); err != nil {
				log.Fatal(err)
			}
		} else {
			// parse dari form
			getID := r.PostFormValue("id")
			id, _ := strconv.Atoi(getID)
			getNim := r.PostFormValue("nim")
			nim, _ := strconv.Atoi(getNim)
			name := r.PostFormValue("name")
			kelas := r.PostFormValue("kelas")
			jurusan := r.PostFormValue("jurusan")
			Mhs = Mahasiswa{
				ID:      id,
				NIM:     nim,
				Name:    name,
				Kelas:   kelas,
				Jurusan: jurusan,
			}
		}

		dataMahasiswa, _ := json.Marshal(Mhs) // to byte
		w.Write(dataMahasiswa)                // cetak di browser
		return
	}

	http.Error(w, "hayo mau ngapain", http.StatusNotFound)
	return
}

func main() {
	http.HandleFunc("/post_mahasiswa", PostMahasiswa)
	fmt.Println("server running...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
