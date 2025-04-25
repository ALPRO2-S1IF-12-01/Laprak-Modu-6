package main

//103112400049 Hisyam Nurdiatmoko
import "fmt"

func hitungRataRata(hisyam []int) float64 {
	var total int
	for _, num := range hisyam {
		total += num
	}
	return float64(total) / float64(len(hisyam))
}

func hitungSimpanganBaku(hisyam []int) float64 {
	rata := hitungRataRata(hisyam)
	var kuadratSelisih float64
	for _, num := range hisyam {
		selisih := float64(num) - rata
		kuadratSelisih += selisih * selisih
	}
	return kuadratSelisih / float64(len(hisyam))
}

func hitungFrekuensi(hisyam []int, target int) int {
	count := 0
	for _, num := range hisyam {
		if num == target {
			count++
		}
	}
	return count
}

func tampilkanIndeks(hisyam []int) {
	fmt.Print("Indeks ganjil: ")
	for i := 1; i < len(hisyam); i += 2 {
		fmt.Print(hisyam[i], " ")
	}
	fmt.Println()

	fmt.Print("Indeks genap: ")
	for i := 0; i < len(hisyam); i += 2 {
		fmt.Print(hisyam[i], " ")
	}
	fmt.Println()
}

func main() {
	var jumlahElemen int
	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(&jumlahElemen)
	hisyam := make([]int, jumlahElemen)
	for i := 0; i < jumlahElemen; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i+1)
		fmt.Scan(&hisyam[i])
	}
	fmt.Println("\nIsi array:", hisyam)
	tampilkanIndeks(hisyam)
	fmt.Printf("Rata-rata: %.2f\n", hitungRataRata(hisyam))
	fmt.Printf("Simpangan baku (tanpa akar): %.2f\n", hitungSimpanganBaku(hisyam))
	var angkaUntukFrekuensi int
	fmt.Print("Masukkan angka untuk dihitung frekuensinya: ")
	fmt.Scan(&angkaUntukFrekuensi)
	fmt.Println("Frekuensi angka", angkaUntukFrekuensi, ":", hitungFrekuensi(hisyam, angkaUntukFrekuensi))
}
