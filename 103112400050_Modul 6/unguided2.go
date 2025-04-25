package main

//103112400050
import (
	"fmt"
)

func hitungRataRata(Ahnaf []int) float64 {
	var total int
	for _, nilai := range Ahnaf {
		total += nilai
	}
	return float64(total) / float64(len(Ahnaf))
}

func hitungSimpanganBaku(Ahnaf []int) float64 {
	rata := hitungRataRata(Ahnaf)
	var jumlahSelisihKuadrat float64
	for _, nilai := range Ahnaf {
		selisih := float64(nilai) - rata
		jumlahSelisihKuadrat += selisih * selisih
	}
	return jumlahSelisihKuadrat / float64(len(Ahnaf))
}

func hitungFrekuensi(Ahnaf []int, target int) int {
	count := 0
	for _, nilai := range Ahnaf {
		if nilai == target {
			count++
		}
	}
	return count
}

func tampilkanIndeks(Ahnaf []int) {
	fmt.Print("Nilai indeks ganjil: ")
	for i := 1; i < len(Ahnaf); i += 2 {
		fmt.Print(Ahnaf[i], " ")
	}
	fmt.Println()

	fmt.Print("Nilai indeks genap: ")
	for i := 0; i < len(Ahnaf); i += 2 {
		fmt.Print(Ahnaf[i], " ")
	}
	fmt.Println()
}

func main() {
	var jumlahElemen int
	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(&jumlahElemen)

	Ahnaf := make([]int, jumlahElemen)
	for i := 0; i < jumlahElemen; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i+1)
		fmt.Scan(&Ahnaf[i])
	}

	fmt.Println("\nIsi array:", Ahnaf)
	tampilkanIndeks(Ahnaf)

	fmt.Printf("Rata-rata: %.2f\n", hitungRataRata(Ahnaf))
	fmt.Printf("Simpangan baku (tanpa akar): %.2f\n", hitungSimpanganBaku(Ahnaf))

	var angkaUntukFrekuensi int
	fmt.Print("Masukkan angka untuk dihitung frekuensinya: ")
	fmt.Scan(&angkaUntukFrekuensi)
	fmt.Printf("Frekuensi angka %d: %d\n", angkaUntukFrekuensi, hitungFrekuensi(Ahnaf, angkaUntukFrekuensi))
}
