package main

import "fmt"

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var pemenang []string

	fmt.Print("Klub A: ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scan(&klubB)

	i := 1
	for {
		fmt.Printf("Pertandingan %d: ", i)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			pemenang = append(pemenang, klubA)
			fmt.Println("Hasil:", klubA)
		} else if skorB > skorA {
			pemenang = append(pemenang, klubB)
			fmt.Println("Hasil:", klubB)
		} else {
			pemenang = append(pemenang, "Draw")
			fmt.Println("Hasil: Draw")
		}
		i++
	}

	fmt.Println("\nDaftar klub pemenang:")
	for i, klub := range pemenang {
		fmt.Printf("Hasil %d: %s\n", i+1, klub)
	}
}
