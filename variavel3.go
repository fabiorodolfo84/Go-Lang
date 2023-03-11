package main


import "fmt"


 func main(){

      var x int = 10

      var y int = 20

     fmt.Printf("valores originais: x=%d, y=%d\n", x, y,)

      //troca de valores

     var temp int = x
     x = y
     y = temp


       fmt.Printf("Valores troados: x=%d, y=%d\n", x,y)


}



