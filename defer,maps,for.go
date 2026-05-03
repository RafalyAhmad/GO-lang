package main

import "fmt"

func main() {
buah:=map[string]int{
"Apel":10,
"Mangga":1,
"Pisang":12,
}

defer fmt.Println("Proses selesai..")
fmt.Println(buah["Apel"])

for nama := range buah {
        fmt.Printf("Nama: %s\n", nama)
    }
}