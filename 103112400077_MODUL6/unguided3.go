package main

import "fmt"

func main() {
	var klubA, klubB string
	var poinA, poinB int
	var pemenang []string

	fmt.Print("Klub A: ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scan(&klubB)
	
	pertandingan := 1

	for {
		fmt.Printf("Pertandingan %d (skor %s dan %s): ", pertandingan, klubA, klubB)
		fmt.Scan(&poinA, &poinB)

		if poinA < 0 || poinB < 0 {
			break
		}

		if poinA > poinB {
			pemenang = append(pemenang, klubA)
		} else if poinB > poinA {
			pemenang = append(pemenang, klubB)
		} else {
			pemenang = append(pemenang, "Draw")
		}
		pertandingan++
	}

	fmt.Println("\nHasil pertandingan:")
	for i := 0; i < len(pemenang); i++ {
		fmt.Printf("Hasil %d: %s\n", i+1, pemenang[i])
	}
}
