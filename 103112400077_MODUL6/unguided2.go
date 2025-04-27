package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&n)
	array := make([]int, n)
	fmt.Println("Masukkan elemen-elemen array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemen ke-%d: ", i)
		fmt.Scan(&array[i])
	}
	// a. Menampilkan keseluruhan isi array
	fmt.Println("\na. Isi keseluruhan array:")
	fmt.Println(array)
	for i := 0; i < len(array); i++ {
		fmt.Printf("array[%d] = %d\n", i, array[i])
	}
	// b. Menampilkan elemen dengan indeks ganjil
	fmt.Println("\nb. Elemen dengan indeks ganjil:")
	for i := 1; i < len(array); i += 2 {
		fmt.Printf("array[%d] = %d\n", i, array[i])
	}
	// c. Menampilkan elemen dengan indeks genap
	fmt.Println("\nc. Elemen dengan indeks genap:")
	for i := 0; i < len(array); i += 2 {
		fmt.Printf("array[%d] = %d\n", i, array[i])
	}
	// d. Cetak elemen kelipatan x
	var x int
	fmt.Print("\nd. Masukkan angka untuk kelipatan indeks: ")
	fmt.Scan(&x)
	fmt.Printf("Elemen dengan indeks kelipatan %d:\n", x)
	if x != 0 {
		for i := 0; i < len(array); i++ {
			if i%x == 0 {
				fmt.Printf("array[%d] = %d\n", i, array[i])
			}
		}
	} else {
		fmt.Println("Tidak bisa mencari kelipatan 0.")
	}
	// e. Hapus elemen Array 
	var hapus int
	fmt.Print("\ne. Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&hapus)

	if hapus >= 0 && hapus < len(array) {
		array = append(array[:hapus], array[hapus+1:]...)
		fmt.Println("Array setelah penghapusan:")
		fmt.Println(array)
		for i := 0; i < len(array); i++ {
			fmt.Printf("array[%d] = %d\n", i, array[i])
		}
	} else {
		fmt.Println("Indeks tidak valid.")
	}
	// f. Rata-rata array
	fmt.Println("\nf. Rata-rata elemen array:")
	if len(array) > 0 {
		sum := 0
		for i := 0; i < len(array); i++ {
			sum += array[i]
		}
		rataRata := float64(sum) / float64(len(array))
		fmt.Printf("%.2f\n", rataRata)
		// g. Standar deviasi
		fmt.Println("\ng. Standar deviasi elemen array:")
		total := 0.0
		for i := 0; i < len(array); i++ {
			total += math.Pow(float64(array[i])-rataRata, 2)
		}
		variance := total / float64(len(array))
		sd := math.Sqrt(variance)
		fmt.Printf("%.2f\n", sd)
	} else {
		fmt.Println("Array kosong, tidak bisa menghitung rata-rata dan standar deviasi.")
	}
	// h. Frekuensi dari suatu bilangan
	var cari int
	fmt.Print("\nh. Masukkan bilangan yang ingin dicari frekuensinya: ")
	fmt.Scan(&cari)

	frekuensi := 0
	for i := 0; i < len(array); i++ {
		if array[i] == cari {
			frekuensi++
		}
	}
	fmt.Printf("Frekuensi %d dalam array adalah: %d\n", cari, frekuensi)
}
