package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&n)
	array := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i)
		fmt.Scan(&array[i])
	}
	fmt.Println("\na. Menampilkan semua elemen array:")
	fmt.Println(array)
	// b. Menampilkan elemen dengan indeks ganjil
	fmt.Println("b. Elemen dengan indeks ganjil:")
	for i := 1; i < len(array); i += 2 {
		fmt.Printf("Indeks %d: %d\n", i, array[i])
	}
	// c. Menampilkan elemen dengan indeks genap
	fmt.Println("c. Elemen dengan indeks genap:")
	for i := 0; i < len(array); i += 2 {
		fmt.Printf("Indeks %d: %d\n", i, array[i])
	}
	// d. Menampilkan elemen dengan indeks kelipatan x
	var x int
	fmt.Print("d. Masukkan nilai x (untuk mencari indeks kelipatan x): ")
	fmt.Scan(&x)
	fmt.Printf("   Elemen dengan indeks kelipatan %d:\n", x)
	for i := 0; i < len(array); i++ {
		if i%x == 0 {
			fmt.Printf("Indeks %d: %d\n", i, array[i])
		}
	}
	// e. Menghapus elemen pada indeks tertentu
	var indeksHapus int
	fmt.Print("e. Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&indeksHapus)
	if indeksHapus >= 0 && indeksHapus < len(array) {
		array = append(array[:indeksHapus], array[indeksHapus+1:]...)
		fmt.Println("   Isi array setelah elemen dihapus:")
		fmt.Println(array)
	} else {
		fmt.Println("   Indeks tidak valid!")
	}
	// f. Menampilkan rata-rata
	total := 0
	for _, v := range array {
		total += v
	}
	rataRata := float64(total) / float64(len(array))
	fmt.Printf("f. Rata-rata dari elemen array: %.2f\n", rataRata)
	// g. Menampilkan simpangan baku (standar deviasi)
	var jumlahSelisihKuadrat float64
	for _, v := range array {
		selisih := float64(v) - rataRata
		jumlahSelisihKuadrat += selisih * selisih
	}
	simpanganBaku := math.Sqrt(jumlahSelisihKuadrat / float64(len(array)))
	fmt.Printf("g. Simpangan baku dari elemen array: %.2f\n", simpanganBaku)
	// h. Menampilkan frekuensi kemunculan bilangan tertentu
	var bilanganCari int
	fmt.Print("h. Masukkan bilangan yang ingin dicari frekuensinya: ")
	fmt.Scan(&bilanganCari)
	frekuensi := 0
	for _, v := range array {
		if v == bilanganCari {
			frekuensi++
		}
	}
	fmt.Printf("   Frekuensi kemunculan angka %d: %d kali\n", bilanganCari, frekuensi)
}
