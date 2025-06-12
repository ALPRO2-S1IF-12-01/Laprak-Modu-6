package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Program Operasi Array")
	fmt.Println("===================")
	angka := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println("Array awal:", angka)
	fmt.Print("\nElemen dengan indeks genap: ")
	for i := 1; i < len(angka); i += 2 {
		fmt.Printf("%d ", angka[i])
	}
	fmt.Println()
	fmt.Print("Elemen dengan indeks ganjil: ")
	for i := 0; i < len(angka); i += 2 {
		fmt.Printf("%d ", angka[i])
	}
	fmt.Println()
	kelipatan := 3
	fmt.Printf("\nElemen dengan indeks kelipatan %d: ", kelipatan)
	for i := 0; i < len(angka); i++ {
		if i%kelipatan == 0 {
			fmt.Printf("%d ", angka[i])
		}
	}
	fmt.Println()
	hapus := 4
	if hapus >= 0 && hapus < len(angka) {
		angka = append(angka[:hapus], angka[hapus+1:]...)
		fmt.Printf("\nArray setelah menghapus indeks %d: %v\n", hapus, angka)
	} else {
		fmt.Println("\nError: Indeks tidak valid!")
		return
	}
	total := 0
	for _, v := range angka {
		total += v
	}
	rata := float64(total) / float64(len(angka))
	fmt.Printf("\nRata-rata: %.2f\n", rata)
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
	fmt.Printf("\nFrekuensi angka %d: %d kali\n", cari, frekuensi)
}
