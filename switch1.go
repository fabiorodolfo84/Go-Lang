//Neste exemplo, o programa solicita ao usuário que digite um número de 1 a 7
// em seguida, usa um switch statement para verificar qual dia da semana corresponde ao número digitado.
// Cada caso corresponde a um dia da semana (de segunda a domingo), e o programa exibe uma mensagem informando qual dia da semana corresponde ao número digitado.
// Se o número não corresponder a um dia da semana válido (ou seja, se não for um número de 1 a 7), o programa exibe uma mensagem informando que o número não corresponde a um dia da semana.

package main

import "fmt"

func main() {
    var num int
    fmt.Print("Digite um número de 1 a 7: ")
    fmt.Scan(&num)

    switch num {
    case 1:
        fmt.Println("O número", num, "corresponde a segunda-feira.")
    case 2:
        fmt.Println("O número", num, "corresponde a terça-feira.")
    case 3:
        fmt.Println("O número", num, "corresponde a quarta-feira.")
    case 4:
        fmt.Println("O número", num, "corresponde a quinta-feira.")
    case 5:
        fmt.Println("O número", num, "corresponde a sexta-feira.")
    case 6:
        fmt.Println("O número", num, "corresponde a sábado.")
    case 7:
        fmt.Println("O número", num, "corresponde a domingo.")
    default:
        fmt.Println("O número", num, "não corresponde a um dia da semana.")
    }
}
