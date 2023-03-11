package main

import "fmt"

// Definindo a struct Pessoa com campos nome, idade, email.

type Pessoa struct {

       nome string
       idade int
       email string
  }
      func main() {
   
          //Criando uma nova pessoa
        pessoa1 := Pessoa{"João", 30, "joao@example.com"}
        fmt.Println("Pessoa 1:", pessoa1)


        //Criando uma nova pessoa com valores atribuidos aos campos individualmentes
        pessoa2:= Pessoa {
        nome: "Maria",
        idade: 25,
        email: "maria@example.com",
     }
        fmt.Println("Pessoa 2:", pessoa2)
         

         //Modificando o valor do campo idade da pessoa1

         pessoa1.idade = 31
         fmt.Println("Nova idade de pessoa 1:", pessoa1.idade)


}
