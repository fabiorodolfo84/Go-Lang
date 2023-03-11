package main

import "fmt"

func main() {
    a := true                  // bool
    b := 10                    // int
    c := int8(20)              // int8
    d := int16(30)             // int16
    e := int32(40)             // int32
    f := int64(50)             // int64
    g := rune('x')             // rune (alias for int32)
    h := 3.14                  // float64 (default floating-point type)
    i := float32(2.71828)      // float32
    j := complex(1.2, 3.4)     // complex128 (default complex type)
    k := complex64(5 + 6i)     // complex64
    l := uintptr(0x123456789)  // uintptr
    
    fmt.Printf("a is of type %T\n", a)
    fmt.Printf("b is of type %T\n", b)
    fmt.Printf("c is of type %T\n", c)
    fmt.Printf("d is of type %T\n", d)
    fmt.Printf("e is of type %T\n", e)
    fmt.Printf("f is of type %T\n", f)
    fmt.Printf("g is of type %T\n", g)
    fmt.Printf("h is of type %T\n", h)
    fmt.Printf("i is of type %T\n", i)
    fmt.Printf("j is of type %T\n", j)
    fmt.Printf("k is of type %T\n", k)
    fmt.Printf("l is of type %T\n", l)
}

