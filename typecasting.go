/*type casting é o processo de converter um valor de um tipo de dados 
em outro tipo de dados compativel. é uma operação comum em muitas linguagens de programação, incluindo Go.
*/


package main

import "fmt"

func main() {

    var x int = 5
    var y float64 = float64(x)


      fmt.Printf("x = %d, y = %f\n", x,y)

}
