
//neste exemplo, vamos criar uma função que recebe um array de numeros inteiiros, 
//retorna a soma de todos os seus elementos.



package main

import "fmt"


func somaArray(arr [] int) int {

      soma := 0

     for _, valor := range arr {
        soma += valor
     }
     return soma 
}
    
     func main() {
     arr := []int{1,2,3,4,5}

     fmt.Println(somaArray(arr))
}




