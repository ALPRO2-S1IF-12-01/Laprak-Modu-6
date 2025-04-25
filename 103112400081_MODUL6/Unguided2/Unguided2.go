//RYAN AKEYLA NOVIANTO WIDODO
//103112400081

package main

import (
	"fmt"
	"math"
)

func hitungRataRata(arr []int) float64 {
	jumlah := 0
	for _, angka := range arr {
		jumlah += angka
	}
	return float64(jumlah) / float64(len(arr))
}

func hitungStandarDeviasi(arr []int, rataRata float64) float64 {
	varian := 0.0
	for _, angka := range arr {
		varian += math.Pow(float64(angka)-rataRata, 2)
	}
	varian /= float64(len(arr))
	return math.Sqrt(varian)
}

func hitungFrekuensi(arr []int, target int) int {
	frekuensi := 0
	for _, angka := range arr {
		if angka == target {
			frekuensi++
		}
	}
	return frekuensi
}

func main() {
	var N int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&N)

	array := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i+1)
		fmt.Scan(&array[i])
	}
	fmt.Println("a. Isi dari semua elemen array:", array)

	fmt.Print("b. Elemen dengan indeks ganjil: ")
	for i := 1; i < N; i += 2 {
		fmt.Print(array[i], " ")
	}
	fmt.Println()

	fmt.Print("c. Elemen dengan indeks genap: ")
	for i := 0; i < N; i += 2 {
		fmt.Print(array[i], " ")
	}
	fmt.Println()

	var x int
	fmt.Print("d. Masukkan bilangan x untuk indeks kelipatan: ")
	fmt.Scan(&x)
	fmt.Print("Indeks kelipatan", x, ": ")
	for i := 0; i < N; i++ {
		if i%x == 0 {
			fmt.Print(array[i], " ")
		}
	}
	fmt.Println()

	var indexToDelete int
	fmt.Print("e. Masukkan indeks untuk dihapus (mulai dari 0): ")
	fmt.Scan(&indexToDelete)
	if indexToDelete >= 0 && indexToDelete < len(array) {
		array = append(array[:indexToDelete], array[indexToDelete+1:]...)
		fmt.Println("Isi array setelah elemen dihapus:", array)
	} else {
		fmt.Println("Indeks di luar battas array.")
	}

	rataRata := hitungRataRata(array)
	fmt.Printf("f. Rata-rata dari bilangan yang ada: %2.f\n", rataRata)

	StandarDeviasi := hitungStandarDeviasi(array, rataRata)
	fmt.Printf("g. Standar Deviasi atau simpangan baku: %2.f\n", StandarDeviasi)

	var target int
	fmt.Print("h. Masukkan bilangan untuk dicari frekuensinya: ")
	fmt.Scan(&target)
	frekuensi := hitungFrekuensi(array, target)
	fmt.Printf("Frekuensi dari bilangan %d: %d\n", target, frekuensi)
}
