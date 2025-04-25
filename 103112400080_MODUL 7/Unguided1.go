// JESIKA METANIA RAHMA ARIFIN
// 10311240080
package main

import "fmt"

func jarakKuadrat(titik1, titik2 [2]int) int {
	return (titik2[0]-titik1[0]) * (titik2[0]-titik1[0]) + (titik2[1]-titik1[1]) * (titik2[1]-titik1[1])

}
func posisiTitik(titik [2]int, lingkaran1, lingkaran2 [3]int) string {
	jarakKuadratKeLingkaran1 := jarakKuadrat(titik, [2]int{lingkaran1[0], lingkaran1[1]})
	jarakKuadratKeLingkaran2 := jarakKuadrat(titik, [2]int{lingkaran2[0], lingkaran2[1]})

	// Kuadrat radius lingkaran
	kuadratRadiusLingkaran1 := lingkaran1[2] * lingkaran1[2]
	kuadratRadiusLingkaran2 := lingkaran2[2] * lingkaran2[2]

	// Tentukan posisi titik di lingkaran 1 dan 2
	diDalamLingkaran1 := jarakKuadratKeLingkaran1 <= kuadratRadiusLingkaran1
	diDalamLingkaran2 := jarakKuadratKeLingkaran2 <= kuadratRadiusLingkaran2

	// Menentukan posisi berdasarkan dua lingkaran
	if diDalamLingkaran1 && diDalamLingkaran2 {
		return "Titik di dalam lingkaran 1 dan 2"
	} else if diDalamLingkaran1 {
		return "Titik di dalam lingkaran 1"
	} else if diDalamLingkaran2 {
		return "Titik di dalam lingkaran 2"
	} else {
		return "Titik di luar lingkaran 1 dan 2"
	}
}

func main() {
	// Input untuk lingkaran 1 (koordinat pusat x, y dan radius)
	var lingkaran1 [3]int
	fmt.Scan(&lingkaran1[0], &lingkaran1[1], &lingkaran1[2])

	// Input untuk lingkaran 2 (koordinat pusat x, y dan radius)
	var lingkaran2 [3]int
	fmt.Scan(&lingkaran2[0], &lingkaran2[1], &lingkaran2[2])

	// Input untuk titik sembarang (koordinat x, y)
	var titik [2]int
	fmt.Scan(&titik[0], &titik[1])

	// Output hasil posisi titik
	fmt.Println(posisiTitik(titik, lingkaran1, lingkaran2))
}