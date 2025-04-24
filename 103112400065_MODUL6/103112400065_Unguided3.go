package main

import (
	"fmt"
)

func main() {
	var club1, club2 string
	var skor1, skor2 int
	var pemenang []string
	DimasRamadhani_103112400065 := 1
	fmt.Print("Klub A: ")
	fmt.Scan(&club1)
	fmt.Print("Klub B: ")
	fmt.Scan(&club2)
	for {
		fmt.Printf("Pertandingan ke-%d : ", DimasRamadhani_103112400065)
		fmt.Scan(&skor1, &skor2)
		if skor1<0 || skor2<0{
			break
		} else if skor1 > skor2 {
			pemenang = append(pemenang, club1)
		} else if skor1 < skor2{
			pemenang = append(pemenang, club2)
		} else {
			pemenang = append(pemenang, "Draw")
		}
		DimasRamadhani_103112400065++
		
	}
	for i := 0; i < len(pemenang); i++ {
		fmt.Printf("Hasil %d: %s\n", i+1, pemenang[i])
	}
	fmt.Print("Pertandingan selesai")
}
