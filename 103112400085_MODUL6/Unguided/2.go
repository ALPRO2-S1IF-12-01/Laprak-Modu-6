//Anastasia Adinda Narendra Indrianto
package main

import "fmt"

func AnastasiaHitungRataRata(arr []int) float64 {
	var total int
	for _, nilai := range arr {
		total += nilai
	}
	return float64(total) / float64(len(arr))
}

func hitungSimpanganBaku(arr []int) float64 {
	rata := AnastasiaHitungRataRata(arr)
	var kuadratSelisih float64
	for _, nilai := range arr {
		selisih := float64(nilai) - rata
		kuadratSelisih += selisih * selisih
	}
	return kuadratSelisih / float64(len(arr))
}

func tampilkanIndeks(arr []int) {
	fmt.Print("Indeks ganjil: ")
	for i := 1; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Print("Indeks genap: ")
	for i := 0; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

func main() {
	var jumlahElemen int
	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(&jumlahElemen)

	arr := make([]int, jumlahElemen)
	for i := 0; i < jumlahElemen; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i+1)
		fmt.Scan(&arr[i])
	}

	fmt.Println("\nIsi array:", arr)
	tampilkanIndeks(arr)
	fmt.Printf("Rata-rata: %.2f\n", AnastasiaHitungRataRata(arr))
	fmt.Printf("Simpangan baku (tanpa akar): %.2f\n", hitungSimpanganBaku(arr))
}
