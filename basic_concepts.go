# 1. Variables:

package main

import "fmt"

func main() {
    // Declare a variable with the `var` keyword
    var a int
    a = 10

    // Declare and initialize a variable with the `:=` shorthand
    b := 20

    fmt.Println(a + b) // 30
}


# 2. Data types:

#package main

import "fmt"

func main() {
    // Integer types
    var a int
    var b int8
    var c int16
    var d int32
    var e int64

    // Unsigned integer types
    var f uint
    var g uint8
    var h uint16
    var i uint32
    var j uint64

    // Floating-point types
    var k float32
    var l float64

    // Boolean type
    var m bool

    // String type
    var n string

    fmt.Printf("%T %T %T %T %T %T %T %T %T %T %T %T %T %T", a, b, c, d, e, f, g, h, i, j, k, l, m, n)
    // Output: int int8 int16 int32 int64 uint uint8 uint16 uint32 uint64 float32 float64 bool string
}


# 3. Loops:

package main

import "fmt"

func main() {
    // For loop
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
    // Output: 0 1 2 3 4

    // While loop
    j := 5
    for j > 0 {
        fmt.Println(j)
        j--
    }
    // Output: 5 4 3 2 1

    // Infinite loop
    for {
        fmt.Println("Infinite loop")
    }
}


# 4. Control Flow:

package main

import "fmt"

func main() {
    // If statement
    a := 10
    if a > 5 {
        fmt.Println("a is greater than 5")
    }

    // If-else statement
    b := 3
    if b > 5 {
        fmt.Println("b is greater than 5")
    } else {
        fmt.Println("b is not greater than 5")
    }

    // Switch statement
    c := "Hello"
    switch c {
    case "Hello":
        fmt.Println("c is Hello")
    case "World":
        fmt.Println("c is World")
    default:
        fmt.Println("c is neither Hello nor World")
    }
}
