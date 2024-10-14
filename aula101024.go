package main

import "fmt"

// Retorna a soma de dois números
func soma(a int, b int) int {
    return a + b
}

// Retorna a soma e o maior número
func somaEmaior(a int, b int) (int, int) {
    maior := a
    if b > a {
        maior = b
    }
    return a + b, maior
}

func main1010() { // Onde o código principal é executado
    fmt.Printf("Hello World!\n")
    
    v := 1
    fmt.Printf("V: %d\n", v)
    
    soma := soma(10, 3)
    fmt.Printf("Soma: %d\n", soma)

    somaMaior, maior := somaEmaior(10, 3)
    fmt.Printf("Soma: %d\n", somaMaior)
    fmt.Printf("Maior: %d\n", maior)
}