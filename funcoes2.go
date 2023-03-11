package main



import (
     "fmt"
     "math"

)


     func isPrime(n int) bool {

        if n <= 1 {

          return false

      }
     for i := 2; i <= int(math.Sqrt(float64(n)));  i++ {
         if n%i == 0 {

             return false


          }

      }
       return true
}


func main() {
     var num int
     fmt.Print("Digite um numero inteiro: ")
     fmt.Scan(&num)
     if isPrime(num) {
       fmt.Println(num, "é um numero primo")
     } else {

          fmt.Println(num, "não é um numero primo")
   }
 }




    
