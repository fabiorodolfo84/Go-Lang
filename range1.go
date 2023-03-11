package main

import "fmt"

func main() {

      palavras := []string{"gato", "cachorro", "papagaio", "coelho"}
      
     for indice, palavra := range palavras {
     fmt.Printf("indice %d: %s\n", indice, palavra)
   }
}



