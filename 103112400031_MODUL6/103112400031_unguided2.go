// Savila Nur Fadilla
// 103112400031

package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Jumlah elemen: ")
	fmt.Scan(&n)

	array := make([]int, n)
	for i := range array {
		fmt.Printf("Elemen [%v]: ", i)
		fmt.Scan(&array[i])
	}

	fmt.Println("\nSemua elemen:", array)

	fmt.Print("Indeks ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Print(array[i], " ")
	}
	fmt.Println()

	fmt.Print("Indeks genap: ")
	for i := 0; i < n; i += 2 {
		fmt.Print(array[i], " ")
	}
	fmt.Println()

	var x int
	fmt.Print("\nKelipatan indeks x: ")
	fmt.Scan(&x)
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(array[i], " ")
		}
	}
	fmt.Println()

	var hapus int
	fmt.Print("\nIndeks yang dihapus: ")
	fmt.Scan(&hapus)
	if hapus >= 0 && hapus < len(array) {
		array = append(array[:hapus], array[hapus+1:]...)
	}
	fmt.Println("array setelah hapus:", array)

	total := 0
	for _, v := range array {
		total += v
	}
	rata := float64(total) / float64(len(array))
	fmt.Printf("\nRata-rata: %.2f\n", rata)

	var deviasi float64
	for _, v := range array {
		deviasi += math.Pow(float64(v)-rata, 2)
	}
	fmt.Printf("Standar deviasi: %.2f\n", math.Sqrt(deviasi/float64(len(array))))

	var cari, frek int
	fmt.Print("\nBilangan yang dicari frekuensinya: ")
	fmt.Scan(&cari)
	for _, v := range array {
		if v == cari {
			frek++
		}
	}
	fmt.Printf("Frekuensi %v: %v kali\n", cari, frek)
}