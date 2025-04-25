package main
// Nufail Alauddin Tsaqif
// 103112400084
import "fmt"
func hitungRataRata(data []int) float64 {
	total := 0
	for _, nilai := range data {
		total += nilai
	}
	return float64(total) / float64(len(data))
}

func hitungSimpanganBaku(data []int) float64 {
	rata := hitungRataRata(data)
	var jumlahKuadrat float64
	for _, nilai := range data {
		selisih := float64(nilai) - rata
		jumlahKuadrat += selisih * selisih
	}
	return jumlahKuadrat / float64(len(data))
}

func hitungFrekuensi(data []int, target int) int {
	hitung := 0
	for _, nilai := range data {
		if nilai == target {
			hitung++
		}
	}
	return hitung
}

func tampilkanBerdasarkanIndeks(data []int, ganjil bool) {
	for i := 0; i < len(data); i++ {
		if ganjil && i%2 != 0 {
			fmt.Print(data[i], " ")
		}
		if !ganjil && i%2 == 0 {
			fmt.Print(data[i], " ")
		}
	}
	fmt.Println()
}

func main() {
	var jumlah int
	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(&jumlah)

	angka := make([]int, jumlah)
	for i := 0; i < jumlah; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i)
		fmt.Scan(&angka[i])
	}

	fmt.Println("\nIsi array:", angka)

	fmt.Print("Elemen pada indeks ganjil: ")
	tampilkanBerdasarkanIndeks(angka, true)

	fmt.Print("Elemen pada indeks genap: ")
	tampilkanBerdasarkanIndeks(angka, false)

	fmt.Printf("Rata-rata: %.2f\n", hitungRataRata(angka))
	fmt.Printf("Simpangan baku (tanpa akar): %.2f\n", hitungSimpanganBaku(angka))

	var cari int
	fmt.Print("Masukkan angka yang ingin dicari frekuensinya: ")
	fmt.Scan(&cari)
	fmt.Printf("Frekuensi angka %d: %d\n", cari, hitungFrekuensi(angka, cari))
}
