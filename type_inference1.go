/*type inference em Go é um recurso que permite ao compilador determinar automaticamente
o tipo de uma variavel com base no valor atribuido a ela. Em outras palavras, o tipo de dados é inferido pelo compilador em vez de ser especificado 
explicitamente pelo programador.
*/

package main


import "fmt"

    func main() {

    x := 5

    y := "Hello, World!"

    fmt.Printf("x is of type %T and y is of type %T\n", x, y)

}



