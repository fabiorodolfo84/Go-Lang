//crie um array de numeros inteiros e calcule a media dos valores contidos no array.

package main

import "fmt"

func main() {
     numeros := [5] int {10, 20, 30, 40, 50}
     
    var soma int
    for _, numero := range numeros {
     soma += numero
} 
     media := float64(soma) / float64(len(numeros))
     
    fmt.Printf("A média dos numeros é: %.2f\n", media)
}



