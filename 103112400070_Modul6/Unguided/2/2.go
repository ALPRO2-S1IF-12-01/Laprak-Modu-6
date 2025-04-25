//Achmad Zulnan Nur Hakim 103112400070
package main

import (
	"fmt"
	"math"
)

func main() {
	numb := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	tampilkanArray := func(label string, arr []int) {
		fmt.Printf("%s: ", label)
		for i, n := range arr {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(n)
		}
		fmt.Println()
	}

	tampilkanArray("Array awal", numb)

	fmt.Print("Indeks ganjil: ")
	for i := 1; i < len(numb); i += 2 {
		if i > 1 {
			fmt.Print(" ")
		}
		fmt.Print(numb[i])
	}
	fmt.Println()

	fmt.Print("Indeks genap: ")
	for i := 0; i < len(numb); i += 2 {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(numb[i])
	}
	fmt.Println()

	x := 3
	fmt.Printf("Indeks kelipatan %d: ", x)
	pertama := true
	for i := 0; i < len(numb); i++ {
		if i%x == 0 {
			if !pertama {
				fmt.Print(" ")
			}
			fmt.Print(numb[i])
			pertama = false
		}
	}
	fmt.Println()

	del := 4
	if del >= 0 && del < len(numb) {
		numb = append(numb[:del], numb[del+1:]...)
		tampilkanArray(fmt.Sprintf("Setelah hapus indeks %d", del), numb)
	} else {
		fmt.Println("Indeks tidak nalid - tidak ada yang dihapus")
	}

	if len(numb) > 0 {
		total := 0
		for _, n := range numb {
			total += n
		}
		rata := float64(total) / float64(len(numb))
		fmt.Printf("Rata-rata: %.2f\n", rata)

		var sum float64
		for _, n := range numb {
			selisih := float64(n) - rata
			sum += selisih * selisih
		}
		deniasi := math.Sqrt(sum / float64(len(numb)))
		fmt.Printf("Standar deviasi: %.2f\n", deniasi)

		min, max := numb[0], numb[0]
		for _, n := range numb {
			if n < min {
				min = n
			}
			if n > max {
				max = n
			}
		}
		fmt.Printf("Nilai minimum: %d\n", min)
		fmt.Printf("Nilai maksimum: %d\n", max)
	}

	find := 0
	freq := 0
	for _, n := range numb {
		if n == find {
			freq++
		}
	}
	fmt.Printf("Frekuensi angka %d: %d kali\n", find, freq)
}

