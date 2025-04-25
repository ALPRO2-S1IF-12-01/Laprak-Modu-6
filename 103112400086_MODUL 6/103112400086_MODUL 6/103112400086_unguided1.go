package main

import "fmt"

type titik struct {
	x, y int
}

type lingkaran struct {
	pusat    titik
	jariJari int
}

func (l lingkaran) berisi(t titik) bool {
	selisihX := t.x - l.pusat.x
	selisihY := t.y - l.pusat.y
	jarakKuadrat := selisihX*selisihX + selisihY*selisihY
	jariJariKuadrat := l.jariJari * l.jariJari
	return jarakKuadrat <= jariJariKuadrat
}

func main() {
	var lingkaranMana [2]lingkaran
	var cekTitik titik

	for i := 0; i < 2; i++ {
		fmt.Printf("masukkan x, y pusat dan jari-jari lingkaran ke-%d : ", i+1)
		fmt.Scan(&lingkaranMana[i].pusat.x, &lingkaranMana[i].pusat.y, &lingkaranMana[i].jariJari)
	}

	fmt.Print("masukkan x dan y titik yang ingin dicek : ")
	fmt.Scan(&cekTitik.x, &cekTitik.y)

	lingkar1 := lingkaranMana[0].berisi(cekTitik)
	lingkar2 := lingkaranMana[1].berisi(cekTitik)

	if lingkar1 && lingkar2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if lingkar1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if lingkar2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}

// CHILA
