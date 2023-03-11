package main

import (
    "fmt"
)

type Aluno struct {
    nome  string
    notas []float64
}

var alunos = make(map[string]Aluno)

func adicionarAluno(nome string, notas []float64) {
    aluno := Aluno{nome, notas}
    alunos[nome] = aluno
}

func exibirNotas(nome string) {
    aluno, ok := alunos[nome]
    if ok {
        fmt.Printf("Notas de %s: %v\n", nome, aluno.notas)
    } else {
        fmt.Printf("Aluno %s não encontrado.\n", nome)
    }
}

func calcularMedia() float64 {
    var total float64
    var count int
    for _, aluno := range alunos {
        for _, nota := range aluno.notas {
            total += nota
            count++
        }
    }
    return total / float64(count)
}

func main() {
    adicionarAluno("João", []float64{8.5, 7.0, 9.0})
    adicionarAluno("Maria", []float64{9.5, 8.0, 9.5})
    adicionarAluno("José", []float64{7.
