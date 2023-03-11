package main

import "fmt"

// neste programa ele recebe uma lista de numeros inteiros
//e retorna a soma dos numeros pares da lista.

func sumEven(numbers[]int) int {

    sum := 0

    for _, num := range numbers {

       if num%2 == 0 {
          sum += num

     }
   }
  
   return sum
}

   func main() {


        numbers := []int{1,2,3,4,5,6,7,8,9,10}
        fmt.Println("A soma dos numeros pares da lista é:", sumEven(numbers))
    }






