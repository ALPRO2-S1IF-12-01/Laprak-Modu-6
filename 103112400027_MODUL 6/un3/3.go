package main

import (
	"fmt"
)

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var pemenang []string
	pertandingan := 1

	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	for {
		fmt.Printf("Pertandingan %d : ", pertandingan)
		_, err := fmt.Scan(&skorA, &skorB)

		if err != nil || skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			fmt.Printf("// %s = %d sedangkan %s = %d\n", klubA, skorA, klubB, skorB)
			fmt.Printf("Hasil %d : %s\n", pertandingan, klubA)
			pemenang = append(pemenang, klubA)
		} else if skorB > skorA {
			fmt.Printf("// %s = %d sedangkan %s = %d\n", klubA, skorA, klubB, skorB)
			fmt.Printf("Hasil %d : %s\n", pertandingan, klubB)
			pemenang = append(pemenang, klubB)
		} else {
			fmt.Printf("// %s = %d sedangkan %s = %d\n", klubA, skorA, klubB, skorB)
			fmt.Printf("Hasil %d : Draw\n", pertandingan)
		}

		pertandingan++
	}

	fmt.Println("Pertandingan selesai")
	fmt.Println("Daftar klub pemenang pertandingan:")

	for i, klub := range pemenang {
		fmt.Printf("Pertandingan %d dimenangkan oleh: %s\n", i+1, klub)
	}
}
