package main
//Mohammad Reyhan Aretha Fatin
// 103112400078
import "fmt"

func rataRata(reyhan []int) float64 {
	var sum int
	for _, v := range reyhan {
		sum += v
	}
	return float64(sum) / float64(len(reyhan))
}

func simpanganBaku(reyhan []int) float64 {
	mean := rataRata(reyhan)
	var sumSquares float64
	for _, v := range reyhan {
		selisih := float64(v) - mean
		sumSquares += selisih * selisih
	}
	return sumSquares / float64(len(reyhan))
}

func frekuensi(reyhan []int, val int) int {
	count := 0
	for _, v := range reyhan {
		if v == val {
			count++
		}
	}
	return count
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(&n)

	reyhan := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i)
		fmt.Scan(&reyhan[i])
	}

	fmt.Println("\nIsi array:", reyhan)

	fmt.Print("Indeks ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Print(reyhan[i], " ")
	}
	fmt.Println()

	fmt.Print("Indeks genap: ")
	for i := 0; i < n; i += 2 {
		fmt.Print(reyhan[i], " ")
	}
	fmt.Println()

	fmt.Printf("Rata-rata: %.2f\n", rataRata(reyhan))
	fmt.Printf("Simpangan baku (tanpa akar): %.2f\n", simpanganBaku(reyhan))

	var x int
	fmt.Print("Masukkan angka untuk dihitung frekuensinya: ")
	fmt.Scan(&x)
	fmt.Println("Frekuensi angka", x, ":", frekuensi(reyhan, x))
}
