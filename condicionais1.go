package main

import "fmt"

 func main() {
      var idade int
      fmt.Print("Qual é a sua idade?")
      fmt.Scan(&idade)

      if idade >= 18 {

           fmt.Println("Você pode dirigir!")

        } else {

            fmt.Println("Você ainda não pode dirigir.")


   } 
}



