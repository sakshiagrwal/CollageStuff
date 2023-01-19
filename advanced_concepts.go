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

package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func (p Person) SayHello() {
    fmt.Printf("Hello, my name is %s and I am %d years old", p.Name, p.Age)
}

func main() {
    p := Person{Name: "John", Age: 30}
    p.SayHello() // Output: "Hello, my name is John and I am 30 years old"
}




// 4. Interfaces:

package main

import "fmt"

type Shape interface {
    Area() float64
}

type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func main() {
    r := Rectangle{Width: 10, Height: 20}
    fmt.Println("Area of rectangle:", r.Area()) // 200

    c := Circle{Radius: 5}
    fmt.Println("Area of circle:", c.Area()) // 78.5
}

