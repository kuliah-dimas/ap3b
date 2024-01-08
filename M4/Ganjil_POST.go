package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Mahasiswa
type Biodata struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Umur   int    `json:"umur"`
	Alamat string `json:"alamat"`
	Nohp   string `json:"no_hp"`
}

// PostMahasiswa
func PostBiodata(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Bio Biodata

	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&Bio); err != nil {
				log.Fatal(err)
			}
		} else {
			// parse dari form
			getID := r.PostFormValue("id")
			id, _ := strconv.Atoi(getID)
			name := r.PostFormValue("name")
			getUmur := r.PostFormValue("umur")
			umur, _ := strconv.Atoi(getUmur)
			alamat := r.PostFormValue("alamat")
			nohp := r.PostFormValue("no_hp")
			Bio = Biodata{
				ID:     id,
				Name:   name,
				Umur:   umur,
				Alamat: alamat,
				Nohp:   nohp,
			}
		}

		databiodata, _ := json.Marshal(Bio) // to byte
		w.Write(databiodata)                // cetak di browser
		return
	}

	http.Error(w, "hayo mau ngapain", http.StatusNotFound)
	return
}

func main() {
	http.HandleFunc("/post_biodata", PostBiodata)
	fmt.Println("server running...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
