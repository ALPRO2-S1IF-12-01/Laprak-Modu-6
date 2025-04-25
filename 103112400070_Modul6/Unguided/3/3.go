//Achmad Zulvan Nur Hakim 103112400070
package main

import "fmt"

func main() {
    var klubA, klubB string
    
    fmt.Print("Klub A : ")
    fmt.Scanln(&klubA)
    
    fmt.Print("Klub B : ")
    fmt.Scanln(&klubB)
    
    fmt.Println()

    var skorA, skorB []int
    pertandinganKe := 1

    for {
        var a, b int
        fmt.Printf("Pertandingan %d : ", pertandinganKe)
        _, err := fmt.Scanln(&a, &b)
        
        if err != nil || a < 0 || b < 0 {
            break
        }
        
        skorA = append(skorA, a)
        skorB = append(skorB, b)
        pertandinganKe++
    }

    for i := 0; i < len(skorA); i++ {
        if skorA[i] > skorB[i] {
            fmt.Printf("Hasil %d : %s\n", i+1, klubA)
        } else if skorB[i] > skorA[i] {
            fmt.Printf("Hasil %d : %s\n", i+1, klubB)
        } else {
            fmt.Printf("Hasil %d : Draw\n", i+1)
        }
    }
    
    fmt.Println("Pertandingan selesai")
}