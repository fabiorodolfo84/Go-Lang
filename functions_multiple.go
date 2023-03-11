package main

import "fmt"

// função que soma dois números inteiros
func somar(a, b int) int {
    return a + b
}

// função que subtrai dois números inteiros
func subtrair(a, b int) int {
    return a - b
}

// função que multiplica dois números inteiros
func multiplicar(a, b int) int {
    return a * b
}

// função que divide dois números inteiros
func dividir(a, b int) float64 {
    return float64(a) / float64(b)
}

func main() {
    // chamando as funções e imprimindo os resultados
    fmt.Println("Resultado da soma: ", somar(5, 3))
    fmt.Println("Resultado da subtração: ", subtrair(10, 4))
    fmt.Println("Resultado da multiplicação: ", multiplicar(6, 7))
    fmt.Println("Resultado da divisão: ", dividir(15, 3))
}
