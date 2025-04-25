package main
//Setyo Nugroho
//103112400024

import "fmt"

func hitungRataRata(kaen []int) float64 {
	var total int
	for _, num := range kaen {
		total += num
	}
	return float64(total) / float64(len(kaen))
}

func hitungSimpanganBaku(kaen []int) float64 {
	rata := hitungRataRata(kaen)
	var kuadratSelisih float64
	for _, num := range kaen {
		selisih := float64(num) - rata
		kuadratSelisih += selisih * selisih
	}
	return kuadratSelisih / float64(len(kaen))
}

func hitungFrekuensi(kaen []int, target int) int {
	count := 0
	for _, num := range kaen {
		if num == target {
			count++
		}
	}
	return count
}

func tampilkanIndeks(kaen []int) {
	fmt.Print("Indeks ganjil: ")
	for i := 1; i < len(kaen); i += 2 {
		fmt.Print(kaen[i], " ")
	}
	fmt.Println()

	fmt.Print("Indeks genap: ")
	for i := 0; i < len(kaen); i += 2 {
		fmt.Print(kaen[i], " ")
	}
	fmt.Println()
}

func main() {
	var jumlahElemen int
	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(&jumlahElemen)
	kaen := make([]int, jumlahElemen)
	for i := 0; i < jumlahElemen; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i+1)
		fmt.Scan(&kaen[i])
	}
	fmt.Println("\nIsi array:", kaen)
	tampilkanIndeks(kaen)
	fmt.Printf("Rata-rata: %.2f\n", hitungRataRata(kaen))
	fmt.Printf("Simpangan baku (tanpa akar): %.2f\n", hitungSimpanganBaku(kaen))
	var angkaUntukFrekuensi int
	fmt.Print("Masukkan angka untuk dihitung frekuensinya: ")
	fmt.Scan(&angkaUntukFrekuensi)
	fmt.Println("Frekuensi angka", angkaUntukFrekuensi, ":", hitungFrekuensi(kaen, angkaUntukFrekuensi))
}