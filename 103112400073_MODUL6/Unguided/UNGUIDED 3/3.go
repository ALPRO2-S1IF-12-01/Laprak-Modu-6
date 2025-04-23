//Muhammad Zaky Mubarok
package main

import "fmt"

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var pemenang []string

	fmt.Print("Klub A : ")
	fmt.Scanln(&klubA)

	fmt.Print("Klub B : ")
	fmt.Scanln(&klubB)

	pertandingan := 1
	for {
		fmt.Printf("Pertandingan %d : ", pertandingan)
		fmt.Scan(&skorA, &skorB)

		// Hentikan jika skor tidak valid (negatif)
		if skorA < 0 || skorB < 0 {
			break
		}

		// Menentukan pemenang atau draw
		if skorA > skorB {
			pemenang = append(pemenang, klubA)
		} else if skorB > skorA {
			pemenang = append(pemenang, klubB)
		} else {
			pemenang = append(pemenang, "Draw")
		}
		pertandingan++
	}

	// Menampilkan hasil pertandingan
	for i, hasil := range pemenang {
		fmt.Printf("Hasil %d : %s\n", i+1, hasil)
	}

	fmt.Println("Pertandingan selesai")
}
