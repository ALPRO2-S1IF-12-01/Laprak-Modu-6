// Felix Pedrosa V 

package main

import (
	"fmt"
	"math"
)

func main() {
	var jumlahElemen int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&jumlahElemen)
	dataArray := make([]int, jumlahElemen)

	// Input elemen array
	for i := 0; i < jumlahElemen; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i)
		fmt.Scan(&dataArray[i])
	}

	// a. Menampilkan semua elemen array
	fmt.Println("\na. Menampilkan semua elemen array:")
	fmt.Println(dataArray)

	// b. Menampilkan elemen dengan indeks ganjil
	fmt.Println("b. Elemen dengan indeks ganjil:")
	for i := 1; i < len(dataArray); i += 2 {
		fmt.Printf("Indeks %d: %d\n", i, dataArray[i])
	}

	// c. Menampilkan elemen dengan indeks genap
	fmt.Println("c. Elemen dengan indeks genap:")
	for i := 0; i < len(dataArray); i += 2 {
		fmt.Printf("Indeks %d: %d\n", i, dataArray[i])
	}

	// d. Menampilkan elemen dengan indeks kelipatan x
	var kelipatan int
	fmt.Print("d. Masukkan nilai x (untuk mencari indeks kelipatan x): ")
	fmt.Scan(&kelipatan)
	fmt.Printf("   Elemen dengan indeks kelipatan %d:\n", kelipatan)
	for i := 0; i < len(dataArray); i++ {
		if i%kelipatan == 0 {
			fmt.Printf("Indeks %d: %d\n", i, dataArray[i])
		}
	}

	// e. Menghapus elemen pada indeks tertentu
	var indeksUntukHapus int
	fmt.Print("e. Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&indeksUntukHapus)
	if indeksUntukHapus >= 0 && indeksUntukHapus < len(dataArray) {
		dataArray = append(dataArray[:indeksUntukHapus], dataArray[indeksUntukHapus+1:]...)
		fmt.Println("   Isi array setelah elemen dihapus:")
		fmt.Println(dataArray)
	} else {
		fmt.Println("   Indeks tidak valid!")
	}

	// f. Menampilkan rata-rata
	totalNilai := 0
	for _, nilai := range dataArray {
		totalNilai += nilai
	}
	rataRata := float64(totalNilai) / float64(len(dataArray))
	fmt.Printf("f. Rata-rata dari elemen array: %.2f\n", rataRata)

	// g. Menampilkan simpangan baku (standar deviasi)
	var totalSelisihKuadrat float64
	for _, nilai := range dataArray {
		selisih := float64(nilai) - rataRata
		totalSelisihKuadrat += selisih * selisih
	}
	simpanganBaku := math.Sqrt(totalSelisihKuadrat / float64(len(dataArray)))
	fmt.Printf("g. Simpangan baku dari elemen array: %.2f\n", simpanganBaku)

	// h. Menampilkan frekuensi kemunculan bilangan tertentu
	var bilanganDicari int
	fmt.Print("h. Masukkan bilangan yang ingin dicari frekuensinya: ")
	fmt.Scan(&bilanganDicari)
	frekuensiKemunculan := 0
	for _, nilai := range dataArray {
		if nilai == bilanganDicari {
			frekuensiKemunculan++
		}
	}
	fmt.Printf("   Frekuensi kemunculan angka %d: %d kali\n", bilanganDicari, frekuensiKemunculan)
}