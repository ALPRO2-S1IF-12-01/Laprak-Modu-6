//M. DAVI ILYAS RENALDO_103112400062
package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(&n)

	data := make([]int, n)
	fmt.Println("Masukkan elemen:")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemen ke-%d: ", i+1)
		fmt.Scan(&data[i])
	}

	fmt.Print("a. Isi: ")
	fmt.Println(data)

	fmt.Print("b. Indeks Ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Print(data[i], " ")
	}
	fmt.Println()

	fmt.Print("c. Indeks Genap: ")
	for i := 0; i < n; i += 2 {
		fmt.Print(data[i], " ")
	}
	fmt.Println()

	var x int
	fmt.Print("d. Masukkan x: ")
	fmt.Scan(&x)
	fmt.Printf("   Kelipatan %d: ", x)
	for i := 0; i < n; i += x {
		fmt.Print(data[i], " ")
	}
	fmt.Println()

	var k int
	fmt.Print("Masukkan k: ")
	fmt.Scan(&k)
	if k >= 0 && k < len(data) {
		data = append(data[:k], data[k+1:]...)
		fmt.Print("   Setelah Hapus: ")
		fmt.Println(data)
	} else {
		fmt.Println("   k tidak valid")
	}

	total := 0.0
	for _, val := range data {
		total += float64(val)
	}
	rata := total / float64(len(data))
	fmt.Printf("e. Rata-rata: %.2f\n", rata)

	jumlahKuadratSelisih := 0.0
	for _, val := range data {
		selisih := float64(val) - rata
		jumlahKuadratSelisih += math.Pow(selisih, 2)
	}
	standarDeviasi := math.Sqrt(jumlahKuadratSelisih / float64(len(data)))
	fmt.Printf("f. Standar Deviasi: %.2f\n", standarDeviasi)

	var y int
	fmt.Print("Masukkan y: ")
	fmt.Scan(&y)
	frekuensi := 0
	for _, val := range data {
		if val == y {
			frekuensi++
		}
	}
	fmt.Printf("g. Frekuensi %d: %d\n", y, frekuensi)
}