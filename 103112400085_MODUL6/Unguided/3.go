//Anastasia Adinda Narendra Indrianto
package main

import "fmt"

func main() {
	var AnastasiaTimSatu, timDua string
	var skorTimSatu, skorTimDua int
	var hasilPertandingan []string

	fmt.Print("Masukkan nama Tim 1: ")
	fmt.Scanln(&AnastasiaTimSatu)
	fmt.Print("Masukkan nama Tim 2: ")
	fmt.Scanln(&timDua)

	hitungTanding := 1
	for {
		fmt.Printf("Masukkan skor untuk pertandingan %d (Tim 1 vs Tim 2): ", hitungTanding)
		fmt.Scan(&skorTimSatu, &skorTimDua)

		if skorTimSatu < 0 || skorTimDua < 0 {
			break
		}

		if skorTimSatu > skorTimDua {
			hasilPertandingan = append(hasilPertandingan, AnastasiaTimSatu)
		} else if skorTimDua > skorTimSatu {
			hasilPertandingan = append(hasilPertandingan, timDua)
		} else {
			hasilPertandingan = append(hasilPertandingan, "Seri")
		}
		hitungTanding++
	}

	for i, hasil := range hasilPertandingan {
		fmt.Printf("Hasil pertandingan %d: %s\n", i+1, hasil)
	}
	fmt.Println("Pertandingan selesai")
}
