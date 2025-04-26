package main

// Muhammad Faris Rachmadi
// 103112400079

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(&N)

	arr := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i+1)
		fmt.Scan(&arr[i])
	}

	fmt.Println("\nIsi array:", arr)

	fmt.Print("Indeks ganjil: ")
	for i := 1; i < N; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Print("Indeks genap: ")
	for i := 0; i < N; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	var sum int
	for i := 0; i < N; i++ {
		sum += arr[i]
	}
	rataRata := float64(sum) / float64(N)
	fmt.Printf("Rata-rata: %.2f\n", rataRata)

	var squaredDiffSum float64
	for i := 0; i < N; i++ {
		squaredDiffSum += math.Pow(float64(arr[i])-rataRata, 2)
	}
	simpanganBakuTanpaAkar := squaredDiffSum / float64(N)
	fmt.Printf("Simpangan baku (tanpa akar): %.2f\n", simpanganBakuTanpaAkar)

	var angka int
	fmt.Print("Masukkan angka untuk dihitung frekuensinya: ")
	fmt.Scan(&angka)
	frekuensi := 0
	for i := 0; i < N; i++ {
		if arr[i] == angka {
			frekuensi++
		}
	}
	fmt.Printf("Frekuensi angka %d: %d\n", angka, frekuensi)
}
