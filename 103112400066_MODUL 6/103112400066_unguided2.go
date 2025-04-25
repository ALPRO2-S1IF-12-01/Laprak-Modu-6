// DWI OKTA SURYANINGRUM
// 103112400066

package main

import (
	"fmt"
	"math"
)

// Struct ArrayHandler menampung array dan jumlah elemen aktif
type ArrayHandler struct {
	data   [100]int // Array berkapasitas maksimal 100 elemen
	length int      // Menyimpan jumlah elemen yang sedang aktif
}

// Menampilkan seluruh isi array
func (a *ArrayHandler) TampilkanSemua() {
	fmt.Println("Isi array:")
	for i := 0; i < a.length; i++ {
		fmt.Printf("%d ", a.data[i]) // Cetak setiap elemen array
	}
	fmt.Println()
}

// Menampilkan elemen dengan indeks ganjil
func (a *ArrayHandler) TampilkanIndeksGanjil() {
	fmt.Println("Elemen dengan indeks ganjil:")
	for i := 1; i < a.length; i += 2 { // Mulai dari indeks 1, lompat 2
		fmt.Printf("data[%d] = %d\n", i, a.data[i])
	}
}

// Menampilkan elemen dengan indeks genap
func (a *ArrayHandler) TampilkanIndeksGenap() {
	fmt.Println("Elemen dengan indeks genap:")
	for i := 0; i < a.length; i += 2 { // Mulai dari indeks 0, lompat 2
		fmt.Printf("data[%d] = %d\n", i, a.data[i])
	}
}

// Menampilkan elemen dengan indeks kelipatan x
func (a *ArrayHandler) TampilkanIndeksKelipatan(x int) {
	fmt.Printf("Elemen dengan indeks kelipatan %d:\n", x)
	for i := 0; i < a.length; i++ {
		if i%x == 0 { // Cek apakah indeks merupakan kelipatan x
			fmt.Printf("data[%d] = %d\n", i, a.data[i])
		}
	}
}

// Menghapus elemen pada indeks tertentu
func (a *ArrayHandler) HapusPadaIndeks(indeks int) {
	for i := indeks; i < a.length-1; i++ {
		a.data[i] = a.data[i+1] // Geser elemen ke kiri
	}
	a.length-- // Kurangi panjang array
	fmt.Println("Setelah penghapusan:")
	a.TampilkanSemua() // Tampilkan isi array setelah elemen dihapus
}

// Menghitung rata-rata
func (a *ArrayHandler) RataRata() float64 {
	total := 0
	for i := 0; i < a.length; i++ {
		total += a.data[i] // Jumlahkan seluruh elemen
	}
	return float64(total) / float64(a.length) // Hitung rata-rata
}

// Menghitung simpangan baku (standar deviasi)
func (a *ArrayHandler) SimpanganBaku() float64 {
	rata := a.RataRata() // Hitung rata-rata dulu
	var total float64
	for i := 0; i < a.length; i++ {
		total += math.Pow(float64(a.data[i])-rata, 2) // (x - rata)^2
	}
	return math.Sqrt(total / float64(a.length)) // Akar dari total dibagi jumlah data
}

// Menampilkan frekuensi dari suatu bilangan
func (a *ArrayHandler) Frekuensi(nilai int) int {
	count := 0
	for i := 0; i < a.length; i++ {
		if a.data[i] == nilai { // Jika elemen sama dengan nilai yang dicari
			count++
		}
	}
	return count // Kembalikan jumlah kemunculan
}

func main() {
	var n int
	handler := ArrayHandler{} // Inisialisasi struct ArrayHandler

	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&n)

	// Input elemen array dari pengguna
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i)
		fmt.Scan(&handler.data[i])
		handler.length++ // Tambah panjang array setiap input
	}

	handler.TampilkanSemua()         // a. Tampilkan semua isi array
	handler.TampilkanIndeksGanjil()  // b. Tampilkan indeks ganjil
	handler.TampilkanIndeksGenap()   // c. Tampilkan indeks genap

	var x int
	fmt.Print("Masukkan nilai x untuk indeks kelipatan: ")
	fmt.Scan(&x)
	handler.TampilkanIndeksKelipatan(x) // d. Tampilkan indeks kelipatan x

	var indeks int
	fmt.Print("Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&indeks)
	handler.HapusPadaIndeks(indeks) // e. Hapus elemen di indeks tertentu

	fmt.Printf("Rata-rata: %.2f\n", handler.RataRata())         // f. Hitung rata-rata
	fmt.Printf("Simpangan baku: %.2f\n", handler.SimpanganBaku()) // g. Hitung standar deviasi

	var cari int
	fmt.Print("Masukkan bilangan untuk dicari frekuensinya: ")
	fmt.Scan(&cari)
	fmt.Printf("Frekuensi %d = %d\n", cari, handler.Frekuensi(cari)) // h. Hitung frekuensi nilai
}