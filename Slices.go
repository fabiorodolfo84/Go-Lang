package main

import (
    "fmt"
    "strconv"
    "strings"
)

func maiorValor(numeros []int) int {
    maior := numeros[0]
    for _, numero := range numeros {
        if numero > maior {
            maior = numero
        }
    }
    return maior
}

func main() {
    // Solicita que o usuário digite um conjunto de números separados por espaço
    fmt.Print("Digite um conjunto de números separados por espaço: ")
    var input string
    fmt.Scanln(&input)

    // Converte a string de entrada em um slice de números inteiros
    partes := strings.Split(input, " ")
    numeros := make([]int, len(partes))
    for i, parte := range partes {
        numero, err := strconv.Atoi(parte)
        if err != nil {
            panic(err)
        }
        numeros[i] = numero
    }

    // Encontra o maior valor no slice e exibe-o ao usuário
    maior := maiorValor(numeros)
    fmt.Printf("O maior valor é: %d\n", maior)
}
