package main

import "fmt"

func main() {
	var klubA, klubB string
	fmt.Println("MULIA AKBAR NANDA PRATAMA\nNIM: 103112400034")
	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	var hasil []string
	pertandingan := 1

	for {
		var skorA, skorB int

		fmt.Printf("Pertandingan %d : ", pertandingan)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}
		if skorA > skorB {
			hasil = append(hasil, klubA)
		} else if skorB > skorA {
			hasil = append(hasil, klubB)
		} else {
			hasil = append(hasil, "Draw")
		}

		pertandingan++
	}

	fmt.Printf("\nHasil 1 : %s\n", klubA)
	fmt.Printf("Hasil 2 : %s\n", klubB)

	for i := 0; i < len(hasil); i++ {
		fmt.Printf("Hasil %d : %s\n", i+3, hasil[i])
	}

	fmt.Println("Pertandingan selesai")
}
