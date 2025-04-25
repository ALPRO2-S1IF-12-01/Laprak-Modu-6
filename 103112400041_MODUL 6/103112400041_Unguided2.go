//BERTHA ADELA
//103112400041
package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah array: ")
	fmt.Scan(&n)

	data := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i+1)
		fmt.Scan(&data[i])
	}

	fmt.Println("\nSeluruh isi array:")
	for i := 0; i < len(data); i++ {
		fmt.Printf("indeks[%d] = %d\n", i+1, data[i])
	}

	fmt.Println("\nElemen dengan indeks ganjil:")
	for i := 0; i < len(data); i++ {
		if (i+1)%2 == 1 {
			fmt.Printf("indeks[%d] = %d\n", i+1, data[i])
		}
	}

	fmt.Println("\nElemen dengan indeks genap:")
	for i := 0; i < len(data); i++ {
		if (i+1)%2 == 0 {
			fmt.Printf("indeks[%d] = %d\n", i+1, data[i])
		}
	}

	var x int
	fmt.Print("\nMasukkan nilai x untuk mencari indeks kelipatan x: ")
	fmt.Scan(&x)
	fmt.Printf("Elemen dengan indeks kelipatan %d:\n", x)
	for i := 0; i < len(data); i++ {
		if x != 0 {
			if (i+1)%x == 0 {
				fmt.Printf("indeks[%d] = %d\n", i+1, data[i])
			}
		}
	}

	var hapus int
	fmt.Print("\nMasukkan indeks yang ingin dihapus: ")
	fmt.Scan(&hapus)
	hapus = hapus - 1

	if hapus >= 0 && hapus < len(data) {
		var temp []int
		for i := 0; i < len(data); i++ {
			if i != hapus {
				temp = append(temp, data[i])
			}
		}
		data = temp
		fmt.Println("Isi array setelah dihapus:")
		for i := 0; i < len(data); i++ {
			fmt.Printf("indeks[%d] = %d\n", i+1, data[i])
		}
	} else {
		fmt.Println("Indeks tidak valid.")
	}

	var jumlah int
	for i := 0; i < len(data); i++ {
		jumlah += data[i]
	}
	var rata2 float64 = float64(jumlah) / float64(len(data))
	fmt.Printf("\nRata-rata: %.2f\n", rata2)

	var total float64
	for i := 0; i < len(data); i++ {
		selisih := float64(data[i]) - rata2
		total += selisih * selisih
	}
	var deviasi float64 = math.Sqrt(total / float64(len(data)))
	fmt.Printf("Standar Deviasi: %.2f\n", deviasi)

	var cari int
	fmt.Print("\nMasukkan angka yang ingin dicari frekuensinya: ")
	fmt.Scan(&cari)
	var frek int
	for i := 0; i < len(data); i++ {
		if data[i] == cari {
			frek++
		}
	}
	fmt.Printf("Frekuensi angka %d: %d kali\n", cari, frek)
}
