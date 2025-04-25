//BERTHA ADELA
//103112400041
package main

import "fmt"

func main() {
    var clubA, clubB string
    var score1, score2 int

    fmt.Print("Klub A: ")
    fmt.Scanln(&clubA)
    fmt.Print("Klub B: ")
    fmt.Scanln(&clubB)

    match := 1
    var results []string

    for score1 != -1 {
        fmt.Printf("Pertandingan ke-%d : ", match)
        fmt.Scanln(&score1, &score2)

		if score1 != -1 {
			if score1 > score2 {
				results = append(results, fmt.Sprintf("Hasil %d : %s", match, clubA))
			} else if score2 > score1 {
				results = append(results, fmt.Sprintf("Hasil %d : %s", match, clubB))
			} else {
				results = append(results, fmt.Sprintf("Hasil %d : Draw", match))
			}
			match = match + 1
		}
    }

    for _, result := range results {
        fmt.Println(result)
    }

    fmt.Println("Pertandingan selesai")
}