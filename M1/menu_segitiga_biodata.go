package main

import (
	"fmt"
)

func main() {
	var (
		pil, uts, uas, tinggi            int
		total                            float64
		nama, npm, kelas, matkul, status string
	)

	fmt.Println("============== Menu pILIHAN============")
	fmt.Println("1. Biodata Nilai")
	fmt.Println("2. Pola bangun datar")
	fmt.Println("Maukkan Pilihan [1-2]: ")
	fmt.Scan(&pil)

	switch pil {
	case 1:
		fmt.Println("")

		fmt.Print("Masukkan namas: ")
		fmt.Scan(&nama)

		fmt.Print("Masukkan npm: ")
		fmt.Scan(&npm)
		fmt.Print("Masukkan kelas: ")
		fmt.Scan(&kelas)
		fmt.Print("Masukkan matkul: ")
		fmt.Scan(&matkul)
		fmt.Print("Masukkan uas: ")
		fmt.Scan(&uas)
		fmt.Print("Masukkan uts: ")
		fmt.Scan(&uts)
		total = (float64(uts) + float64(uas)) / 2

		if total >= 90 {
			status = "Lulus, Grade A"
		} else if total >= 80 {
			status = "Lulus, Grade B"
		} else if total >= 70 {
			status = "Lulus, Grade C"
		} else if total >= 60 {
			status = "Lulus, Grade D"
		} else {
			status = "Tidak Lulus, Grade E"
		}

		fmt.Println("==== HASIL ====")
		fmt.Printf("Nama saya %s \n", nama)
		fmt.Printf("Npm saya %s \n", npm)
		fmt.Printf("Kelas saya %s \n", kelas)
		fmt.Printf("Matkul saya %s \n", matkul)
		fmt.Printf("Total saya %3.2f \n", total)
		fmt.Printf("Status: %s \n", status)
	case 2:
		fmt.Print("Masukkan tinggi segitiga: ")
		fmt.Scan(&tinggi)

		fmt.Println("Pola segitiga")
		for i := 1; i <= tinggi; i++ {
			for j := 1; j <= i; j++ {
				fmt.Print("* ")
			}
			fmt.Println("")
		}
	default:
		fmt.Println("Pilihan tidak adas")
	}
}
