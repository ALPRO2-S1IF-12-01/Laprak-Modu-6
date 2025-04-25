package main
//Setyo Nugroho
//103112400024

import "fmt"

const NMAX = 100

func main() {
	var klubA, klubB string
	var hasilPertandingan [NMAX]string
	var skorA, skorB int
	var matchCount int = 0
	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)
	
	for {
		matchCount++
		fmt.Printf("Pertandingan %d : ", matchCount)
		fmt.Scan(&skorA, &skorB)
	
		if skorA < 0 || skorB < 0 {
			matchCount--
			break
		}
		
		if skorA > skorB {
			hasilPertandingan[matchCount-1] = klubA
		} else if skorB > skorA {
			hasilPertandingan[matchCount-1] = klubB
		} else {
			hasilPertandingan[matchCount-1] = "Draw"
		}
	}
	for i := 0; i < matchCount; i++ {
		fmt.Printf("Hasil %d : %s\n", i+1, hasilPertandingan[i])
	}
	
	fmt.Println("Pertandingan selesai")
}