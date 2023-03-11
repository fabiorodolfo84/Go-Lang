package main

import "fmt"

// função que recebe um nome e uma idade e retorna uma mensagem formatada
func saudacao(nome string, idade int) (mensagem string, err error) {
    if idade < 0 {
        err = fmt.Errorf("idade inválida: %d", idade)
        return
    }

    mensagem = fmt.Sprintf("Olá, %s! Você tem %d anos.", nome, idade)
    return
}

func main() {
    // chamando a função saudacao e imprimindo a mensagem retornada
    mensagem, err := saudacao("João", 30)
    if err != nil {
        fmt.Println("Erro ao criar saudação:", err)
        return
    }
    fmt.Println(mensagem)
}
