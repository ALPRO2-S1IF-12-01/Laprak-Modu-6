package main

//103112400049 Hisyam Nurdiatmoko

import "fmt"

func main() {
	var klubA, klubB string
	var skor1, skor2 int
	var pemenang_hisyam []string
	var pertandingan_103112400049 int = 1

	fmt.Print("Klub A: ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scan(&klubB)

	for {
		fmt.Printf("Pertandingan ke-%d : ", pertandingan_103112400049)
		fmt.Scan(&skor1, &skor2)
		if skor1 < 0 || skor2 < 0 {
			break
		}
		if skor1 > skor2 {
			pemenang_hisyam = append(pemenang_hisyam, klubA)
		} else if skor1 < skor2 {
			pemenang_hisyam = append(pemenang_hisyam, klubB)
		} else {
			pemenang_hisyam = append(pemenang_hisyam, "Draw")
		}
		pertandingan_103112400049++
	}
	for i, result := range pemenang_hisyam {
		fmt.Printf("Hasil %d: %s\n", i+1, result)
	}
	fmt.Println("Pertandingan selesai")
}
