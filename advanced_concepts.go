// 1. Functions:

package main

import "fmt"

func add(a int, b int) int {
    return a + b
}

func main() {
    x := 3
    y := 4
    z := add(x, y)
    fmt.Println(z) // 7
}
