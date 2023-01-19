// 1. Functions:


package main

import "fmt"

// Function that takes two integers as arguments and returns their sum
func add(a int, b int) int {
    return a + b
}

// Function that takes a variable number of integers as arguments and returns their sum
func addAll(numbers ...int) int {
    sum := 0
    for _, num := range numbers {
        sum += num
    }
    return sum
}

func main() {
    fmt.Println(add(1, 2)) // 3
    fmt.Println(addAll(1, 2, 3, 4)) // 10
}



// 2. Modules:

// In a file named "mymodule.go"
package mymodule

func Greet() {
    fmt.Println("Hello from mymodule")
}

// In a file named "main.go"
package main

import (
    "fmt"
    "mymodule"
)

func main() {
    mymodule.Greet() // Output: "Hello from mymodule"
}


// 3. Structs:


// 4. Interfaces:
