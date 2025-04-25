package main

import (
	"fmt"
	"math"
)

func main() {
	var n, x, hapusIdx, cari int
	fmt.Print("Jumlah elemen: ")
	fmt.Scan(&n)
	data := make([]int, n)
	fmt.Println("Masukkan elemen:")
	for i := 0; i < n; i++ {
		fmt.Printf("indeks ke-%d: ", i)
		fmt.Scan(&data[i])
	}

	fmt.Print("a. Isi array: ")
	for _, v := range data {
		fmt.Print(v, " ")
	}
	fmt.Println()

	fmt.Print("b. Indeks ganjil: ")
	for i := 1; i < len(data); i += 2 {
		fmt.Print(data[i], " ")
	}
	fmt.Println()

	fmt.Print("c. Indeks genap: ")
	for i := 0; i < len(data); i += 2 {
		fmt.Print(data[i], " ")
	}
	fmt.Println()

	fmt.Print("Masukkan Indeks kelipatan x: ")
	fmt.Scan(&x)
	fmt.Printf("d. Indeks kelipatan %d: ", x)
	for i := 0; i < len(data); i++ {
		if i%x == 0 {
			fmt.Print(data[i], " ")
		}
	}
	fmt.Println()

	fmt.Print("Masukan Indeks yang ingin dihapus: ")
	fmt.Scan(&hapusIdx)
	data = append(data[:hapusIdx], data[hapusIdx+1:]...) 

	fmt.Print("e. Array Setelah dihapus: ")
	for _, v := range data {
		fmt.Print(v, " ")
	}
	fmt.Println()

	var total float64
	for _, v := range data {
		total += float64(v)
	}
	rata := total / float64(len(data))
	fmt.Printf("f. Rata-rata: %.2f\n", rata)

	var jumlahKuadrat float64
	for _, v := range data {
		selisih := float64(v) - rata
		jumlahKuadrat += selisih * selisih
	}
	sd := math.Sqrt(jumlahKuadrat / float64(len(data)))
	fmt.Printf("g. Simpangan baku: %.2f\n", sd)

	fmt.Print("Masukan Nilai yang ingin dicari frekuensinya: ")
	fmt.Scan(&cari)
	freq := 0
	for _, v := range data {
		if v == cari {
			freq++
		}
	}
	fmt.Printf("h. Frekuensi %d: %d\n", cari, freq)
}