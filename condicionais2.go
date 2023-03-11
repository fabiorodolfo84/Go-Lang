package main

import "fmt"

func main() {

   var num int
   fmt.Print("Digite um número: ")

   fmt.Scan(&num)


   if num%3 == 0 && num%5 == 0 {


     fmt.Println("O número é divisivel por 3 e 5.")

     } else if num%3 == 0 {

       fmt.Println("O número é divisivel por 3.")

     } else if num%5 == 0 {

       fmt.Println("O Número é divisivel por 5.")

     } else {

       fmt.Println("O número não é divisivel por 3 ou 5.")
 }

}

