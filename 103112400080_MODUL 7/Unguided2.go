// JESIKA METANIA RAHMA ARIFIN
// 103112400080

package main

import (
	"fmt"
	"math"
)

func inputArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i)
		fmt.Scan(&arr[i])
	}
	return arr
}

func tampilSemua(arr []int) {
	fmt.Println("Isi array:", arr)
}

func tampilIndeksGanjil(arr []int) {
	fmt.Print("Indeks ganjil: ")
	for i := 1; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

func tampilIndeksGenap(arr []int) {
	fmt.Print("Indeks genap: ")
	for i := 0; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

func tampilKelipatan(arr []int, x int) {
	fmt.Printf("Indeks kelipatan %d: ", x)
	for i := 0; i < len(arr); i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()
}

func hapusIndeks(arr []int, idx int) []int {
	if idx < 0 || idx >= len(arr) {
		fmt.Println("Indeks tidak valid!")
		return arr
	}
	return append(arr[:idx], arr[idx+1:]...)
}

func rataRata(arr []int) float64 {
	total := 0
	for _, v := range arr {
		total += v
	}
	return float64(total) / float64(len(arr))
}

func standarDeviasi(arr []int) float64 {
	avg := rataRata(arr)
	var total float64
	for _, v := range arr {
		selisih := float64(v) - avg
		total += selisih * selisih
	}
	return math.Sqrt(total / float64(len(arr)))
}

func frekuensi(arr []int, angka int) int {
	count := 0
	for _, v := range arr {
		if v == angka {
			count++
		}
	}
	return count
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&n)

	arr := inputArray(n)
	tampilSemua(arr)
	tampilIndeksGanjil(arr)
	tampilIndeksGenap(arr)

	var x int
	fmt.Print("Masukkan nilai x untuk kelipatan indeks: ")
	fmt.Scan(&x)
	tampilKelipatan(arr, x)

	var hapus int
	fmt.Print("Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&hapus)
	arr = hapusIndeks(arr, hapus)
	tampilSemua(arr)

	if len(arr) > 0 {
		fmt.Printf("Rata-rata: %.2f\n", rataRata(arr))
		fmt.Printf("Standar deviasi: %.2f\n", standarDeviasi(arr))
	} else {
		fmt.Println("Array kosong, tidak bisa menghitung rata-rata dan standar deviasi.")
	}

	var cari int
	fmt.Print("Masukkan bilangan untuk dihitung frekuensinya: ")
	fmt.Scan(&cari)
	fmt.Printf("Frekuensi %d: %d kali\n", cari, frekuensi(arr, cari))
}
