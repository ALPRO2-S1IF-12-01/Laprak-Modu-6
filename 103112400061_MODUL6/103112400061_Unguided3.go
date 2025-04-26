package main

import "fmt"

func main() {
	var ( 
		match [50]string
		klubA, klubB string
		skorKlubA, skorKlubB, jumlahMatch int
	)

	fmt.Print("Klub A: ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scan(&klubB)

	for i := 0; i < len(match); i++ {
		fmt.Printf("Pertandingan %d : ", i+1)
		fmt.Scan(&skorKlubA, &skorKlubB)
		if skorKlubA < 0 || skorKlubB < 0 {
			break
		}
		if skorKlubA > skorKlubB {
			match[i] = match[i] + klubA
		} else if skorKlubA < skorKlubB {
			match[i] = match[i] + klubB
		} else {
			match[i] = match[i] + "Draw"
		}
		jumlahMatch += 1
	}

	for i := 0; i < jumlahMatch; i++ {
		fmt.Printf("Hasil %d : %s\n", i+1, match[i])
	}
	fmt.Println("Pertandingan selesai")
}