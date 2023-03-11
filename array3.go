//nesse programa criamos um script de como remover um elemento de um array.


package main

import "fmt"


 func removeElemento(arr []int, elemento int) []int {
     novoArray  := []int{}
     for _, valor := range arr {

         if valor != elemento {

             novoArray = append(novoArray, valor)
     }
   }
return novoArray

}

 func main() {

      arr := []int{1,2,3,4,5}

      novoArr := removeElemento(arr, 3)

       fmt.Println(novoArr)

}



