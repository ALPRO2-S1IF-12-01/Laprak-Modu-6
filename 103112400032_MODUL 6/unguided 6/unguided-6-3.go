// daffa tsaqifna f
package main

import "fmt"

func main() {
	var tim1, tim2 string
	var score1, score2, x int
	var result []string
	x = 0
	fmt.Print("nama tim 1: ")
	fmt.Scan(&tim1)
	fmt.Print("nama tim 2: ")
	fmt.Scan(&tim2)
	for {
		fmt.Print("score pertandingan = ")
		_, err := fmt.Scan(&score1, &score2)
		if err != nil || score1 < 0 || score2 < 0 {
			fmt.Println("Pertandingan selesai")
			break
		}
		if score1 > score2 {
			result = append(result, tim1)
		} else if score1 < score2 {
			result = append(result, tim2)
		} else if score1 == score2 {
			result = append(result, "draw")
		}
		x += 1
	}
	for i := 0; i < x; i++ {
		fmt.Printf("Hasil %d : %s\n", i+1, result[i])
	}
}
