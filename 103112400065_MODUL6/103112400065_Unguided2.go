

package main

import (
	"fmt"
	"math"
)

func main() {
	var N, n, hapus, hasil int
	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(&N)
	bilArr := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Printf("Masukkan indeks ke-%d: ", i)
		fmt.Scan(&bilArr[i])
	}

	// soal a
	fmt.Println("=== SOAL A ===")
	fmt.Println("Isi dari array:", bilArr)

	// soal b
	fmt.Println()
	fmt.Println("=== SOAL B ===")
	for i := 0; i < N; i++ {
		if i%2 != 0 {
			fmt.Println("Elemen dengan indeks ganjil:", bilArr[i])
		}
	}

	// soal c
	fmt.Println()
	fmt.Println("=== SOAL C ===")
	for i := 0; i < N; i++ {
		if i%2 == 0 {
			fmt.Println("Elemen dengan indeks genap:", bilArr[i])
		}
	}

	// soal d, menapilkan elemen array dengan indeks keliparan n
	fmt.Println()
	fmt.Println("=== SOAL D ===")
	fmt.Print("Masukkan keliapatan: ")
	fmt.Scan(&n)
	for i := 0; i < len(bilArr); i++ {
		if i%n == 0 && i != 0 {
			fmt.Printf("Elemen dengan indeks kelipatan ke-%d: %d\n", n, bilArr[i])
		}
	}

	// soal e
	fmt.Println()
	fmt.Println("=== SOAL E ===")
	fmt.Print("Menghapus elemen ke-")
	fmt.Scan(&hapus)
	bilArr = append(bilArr[:hapus], bilArr[hapus+1:]...)
	fmt.Printf("Array setelah menghapus elemen ke-%d: %d\n", hapus, bilArr)

	// soal f
	var rataRata float64
	fmt.Println()
	fmt.Println("=== SOAL F ===")
	for _, f := range bilArr {
		hasil += f
	}
	rataRata = float64(hasil) / float64(len(bilArr))
	fmt.Printf("Rata-ratanya dari isi array: %.2f\n", rataRata)

	// soal g
	fmt.Println()
	fmt.Println("=== SOAL G ===")
	var simpBaku, selisih, jumSelisih float64
	for _, xi := range bilArr {
		selisih = float64(xi) - float64(rataRata)
		jumSelisih += math.Pow(selisih, 2)
	}
	simpBaku = math.Sqrt(jumSelisih / float64(len(bilArr)))
	fmt.Printf("Simpangan Baku:%.2f\n", simpBaku)

	fmt.Println("\n===============")
	fmt.Println("Dimas Ramadhani")
	fmt.Println("103112400065")
	fmt.Print("===============\n")
	// Soal h
	var h, frek int
	fmt.Println()
	fmt.Println("=== SOAL H ===")
	fmt.Print("Masukkan angka untuk mencari frekuensinya: ")
	fmt.Scan(&h)
	for _, ha:= range bilArr{
		if h==ha {
			frek++
		}
	}
	fmt.Printf("Jadi jumlah angka %d adalah %d", h, frek)
}
