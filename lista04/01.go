package main

import (
	"fmt"
)

func main() {
	var numeros [10]int
	var encontrados bool

	for i := 0; i < 10; i++ {
		fmt.Scan(&numeros[i])
	}

	for i, valor := range numeros {
		if valor > 50 {
			fmt.Printf("Número: %d, posição: %d\n", valor, i)
			encontrados = true
		}
	}

	if !encontrados {
		fmt.Println("Não há números maiores que 50")
	}
}
