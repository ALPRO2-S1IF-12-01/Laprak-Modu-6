package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("NAMA: MULIA AKBAR NANDA PRATAMA\nNIM: 103112400034")
	angka := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println("Array:", angka)

	fmt.Print("Indeks genap: ")
	for i := 1; i < len(angka); i += 2 {
		fmt.Print(angka[i], " ")
	}
	fmt.Println()

	fmt.Print("Indeks ganjil: ")
	for i := 0; i < len(angka); i += 2 {
		fmt.Print(angka[i], " ")
	}
	fmt.Println()
	
	x := 3
	fmt.Printf("Indeks kelipatan %d: ", x)
	for i := 0; i < len(angka); i++ {
		if i%x == 0 {
			fmt.Print(angka[i], " ")
		}
	}
	fmt.Println()

	hapus := 4
	angka = append(angka[:hapus], angka[hapus+1:]...)
	fmt.Println("Setelah hapus index 4:", angka)

	total := 0
	for _, v := range angka {
		total += v
	}
	rata := float64(total) / float64(len(angka))
	fmt.Printf("Rata-rata: %.2f\n", rata)

	var sum float64
	for _, v := range angka {
		sel := float64(v) - rata
		sum += sel * sel
	}
	stdev := math.Sqrt(sum / float64(len(angka)))
	fmt.Printf("Standar deviasi: %.2f\n", stdev)

	cari := 60
	frekuensi := 0
	for _, v := range angka {
		if v == cari {
			frekuensi++
		}
	}
	fmt.Printf("Frekuensi angka %d: %d kali\n", cari, frekuensi)
}
