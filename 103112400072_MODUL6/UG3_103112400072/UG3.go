package main

import "fmt"

func main() {
	var jumPer int
	fmt.Print("Masukkan jumlah pertandingan: ")
	fmt.Scan(&jumPer)

	var hasil string

	for i := 1; i <= jumPer; i++ {
		var mu, inter int
		fmt.Scan(&mu, &inter)

		fmt.Printf("Pertandingan %d : %d %d\n", i, mu, inter)

		if mu > inter {
			hasil += fmt.Sprintf("Hasil %d : MU\n", i)
		} else if mu < inter {
			hasil += fmt.Sprintf("Hasil %d : Inter\n", i)
		} else {
			hasil += fmt.Sprintf("Hasil %d : Draw\n", i)
		}
	}

	fmt.Print(hasil)
	fmt.Println("Pertandingan selesai")
}
