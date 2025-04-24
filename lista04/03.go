package main

import "fmt"

func main() {
    var numeros [10]int
    var pares []int
    var impares []int
    somaPares := 0

    fmt.Println("Digite 10 números inteiros:")
    for i := 0; i < 10; i++ {
        fmt.Printf("Número %d: ", i+1)
        fmt.Scan(&numeros[i])
        
        if numeros[i]%2 == 0 {
            pares = append(pares, numeros[i])
            somaPares += numeros[i]
        } else {
            impares = append(impares, numeros[i])
        }
    }

    fmt.Println("\na) Números pares digitados:")
    if len(pares) > 0 {
        for _, num := range pares {
            fmt.Printf("%d ", num)
        }
    } else {
        fmt.Println("Nenhum número par foi digitado.")
    }

    fmt.Printf("\n\nb) Soma dos números pares: %d\n", somaPares)

    fmt.Println("\nc) Números ímpares digitados:")
    if len(impares) > 0 {
        for _, num := range impares {
            fmt.Printf("%d ", num)
        }
    } else {
        fmt.Println("Nenhum número ímpar foi digitado.")
    }

    fmt.Printf("\n\nd) Quantidade de números ímpares: %d\n", len(impares))
}