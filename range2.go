package main

import "fmt"

func main() {
    temperaturas := []int{25, 26, 28, 31, 29, 27, 26}
    diasDaSemana := []string{"domingo", "segunda", "terça", "quarta", "quinta", "sexta", "sábado"}

    for i, temperatura := range temperaturas {
        dia := diasDaSemana[i]
        fmt.Printf("%s: %d\n", dia, temperatura)
    }
}

