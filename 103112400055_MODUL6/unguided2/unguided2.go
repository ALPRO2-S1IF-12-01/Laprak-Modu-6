//Feros Pedrosa

package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&N)

	array := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i)
		fmt.Scan(&array[i])
	}

	// a.
	fmt.Println("a. Isi dari semua elemen array:", array)

	// b.
	fmt.Print("b. Elemen dengan indeks ganjil: ")
	for i := 1; i < N; i += 2 {
		fmt.Print(array[i], " ")
	}
	fmt.Println()

	// c.
	fmt.Print("c. Elemen dengan indeks genap: ")
	for i := 0; i < N; i += 2 {
		fmt.Print(array[i], " ")
	}
	fmt.Println()

	// d.
	var x int
	fmt.Print("d. Masukkan bilangan x untuk indeks kelipatan: ")
	fmt.Scan(&x)
	fmt.Print("Indeks kelipatan ", x, ": ")
	for i := 0; i < N; i++ {
		if i%x == 0 {
			fmt.Print(array[i], " ")
		}
	}
	fmt.Println()

	// e.
	var indexToDelete int
	fmt.Print("e. Masukkan indeks untuk dihapus: ")
	fmt.Scan(&indexToDelete)
	array = append(array[:indexToDelete], array[indexToDelete+1:]...)
	fmt.Println("Isi array setelah elemen dihapus:", array)

	// f.
	var sum int
	for _, value := range array {
		sum += value
	}
	average := float64(sum) / float64(len(array))
	fmt.Printf("f. Rata-rata dari bilangan yang ada: %.2f\n", average)

	// g.
	var variance float64
	for _, value := range array {
		variance += math.Pow(float64(value)-average, 2)
	}
	variance /= float64(len(array))
	stdDev := math.Sqrt(variance)
	fmt.Printf("g. Standar Deviasi atau simpangan baku: %.2f\n", stdDev)

	// h.
	var target int
	fmt.Print("h. Masukkan bilangan untuk dicari frekuensinya: ")
	fmt.Scan(&target)
	frequency := 0
	for _, value := range array {
		if value == target {
			frequency++
		}
	}
	fmt.Printf("Frekuensi dari bilangan %d: %d\n", target, frequency)
}
