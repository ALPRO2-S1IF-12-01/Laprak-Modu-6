package main

import "fmt"

func main() {
	var jumlahElemen int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&jumlahElemen)

	data := make([]int, jumlahElemen)
	for i := 0; i < jumlahElemen; i++ {
		fmt.Printf("Data ke-%d: ", i)
		fmt.Scan(&data[i])
	}

	// a. Menampilkan seluruh isi array
	fmt.Println("a. Seluruh isi array:", data)

	// b. Menampilkan elemen dengan indeks ganjil
	fmt.Print("b. Elemen dengan indeks ganjil: ")
	for i := 1; i < len(data); i += 2 {
		fmt.Print(data[i], " ")
	}
	fmt.Println()

	// c. Menampilkan elemen dengan indeks genap
	fmt.Print("c. Elemen dengan indeks genap: ")
	for i := 0; i < len(data); i += 2 {
		fmt.Print(data[i], " ")
	}
	fmt.Println()

	// d. Menampilkan elemen dengan indeks kelipatan x
	var kelipatan int
	fmt.Print("Masukkan nilai x (kelipatan indeks): ")
	fmt.Scan(&kelipatan)
	fmt.Printf("d. Elemen dengan indeks kelipatan %d: ", kelipatan)
	for i := 0; i < len(data); i++ {
		if i%kelipatan == 0 {
			fmt.Print(data[i], " ")
		}
	}
	fmt.Println()

	// e. Menghapus elemen pada indeks tertentu
	var indeksHapus int
	fmt.Print("Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&indeksHapus)
	if indeksHapus >= 0 && indeksHapus < len(data) {
		data = append(data[:indeksHapus], data[indeksHapus+1:]...)
	}
	fmt.Println("e. Array setelah penghapusan:", data)

	// f. Menghitung rata-rata
	var total int
	for _, nilai := range data {
		total += nilai
	}
	rataRata := float64(total) / float64(len(data))
	fmt.Printf("f. Rata-rata: %.2f\n", rataRata)

	// g. Menghitung simpangan baku kuadrat (tanpa akar kuadrat)
	var jumlahSelisihKuadrat float64
	for _, nilai := range data {
		selisih := float64(nilai) - rataRata
		jumlahSelisihKuadrat += selisih * selisih
	}
	fmt.Printf("g. Simpangan baku kuadrat: %.2f\n", jumlahSelisihKuadrat/float64(len(data)))

	// h. Menghitung frekuensi kemunculan angka tertentu
	var angkaCari, frekuensi int
	fmt.Print("Masukkan angka yang ingin dihitung frekuensinya: ")
	fmt.Scan(&angkaCari)
	for _, nilai := range data {
		if nilai == angkaCari {
			frekuensi++
		}
	}
	fmt.Printf("h. Frekuensi kemunculan angka %d: %d kali\n", angkaCari, frekuensi)
}
