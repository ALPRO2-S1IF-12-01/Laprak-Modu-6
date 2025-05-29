package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&N)

	angka := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i)
		fmt.Scan(&angka[i])
	}

	fmt.Println("Isi array:", angka)

	fmt.Print("Indeks ganjil: ")
	for i := 1; i < N; i += 2 {
		fmt.Print(angka[i], " ")
	}
	fmt.Println()

	fmt.Print("Indeks genap: ")
	for i := 0; i < N; i += 2 {
		fmt.Print(angka[i], " ")
	}
	fmt.Println()

	var x int
	fmt.Print("Masukkan bilangan x: ")
	fmt.Scan(&x)
	fmt.Print("Indeks kelipatan ", x, ": ")
	for i := 0; i < N; i++ {
		if i%x == 0 {
			fmt.Print(angka[i], " ")
		}
	}
	fmt.Println()

	var indeksHapus int
	fmt.Print("Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&indeksHapus)
	angka = append(angka[:indeksHapus], angka[indeksHapus+1:]...)
	fmt.Println("Isi array setelah penghapusan:", angka)

	var total int
	for _, v := range angka {
		total += v
	}
	rataRata := float64(total) / float64(len(angka))
	fmt.Printf("Rata-rata: %.2f\n", rataRata)

	var totalKuadrat float64
	for _, v := range angka {
		totalKuadrat += math.Pow(float64(v)-rataRata, 2)
	}
	standarDeviasi := math.Sqrt(totalKuadrat / float64(len(angka)))
	fmt.Printf("Standar deviasi: %.2f\n", standarDeviasi)

	var bilangan int
	fmt.Print("Masukkan bilangan untuk frekuensi: ")
	fmt.Scan(&bilangan)
	frekuensi := 0
	for _, v := range angka {
		if v == bilangan {
			frekuensi++
		}
	}
	fmt.Printf("Frekuensi dari %d: %d\n", bilangan, frekuensi)
}
