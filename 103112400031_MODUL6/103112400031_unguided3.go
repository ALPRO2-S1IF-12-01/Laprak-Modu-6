// Savila Nur Fadilla
// 103112400031

package main

import "fmt"

func main() {
	var klubA, klubB string
	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	var pemenang []string
	pertandingan := 1

	for {
		var skorA, skorB int
		fmt.Printf("Pertandingan %v : ", pertandingan)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			pemenang = append(pemenang, klubA)
		} else if skorB > skorA {
			pemenang = append(pemenang, klubB)
		} else {
			pemenang = append(pemenang, "Draw")
		}
		pertandingan++
	}
	
	fmt.Println()
	for i, hasil := range pemenang {
		fmt.Printf("Hasil %-2d : %s\n", i+1, hasil)
	}

	fmt.Println("Pertandingan selesai")
}
