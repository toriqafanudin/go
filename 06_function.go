package main

import "fmt"

// Fungsi sederhana tanpa parameter dan return value
func sayHello() {
    fmt.Println("Hello, World!")
}

// Fungsi dengan parameter dan return value
func add(x int, y int) int {
    return x + y
}

// Fungsi dengan multiple return value
func calculate(x int, y int) (int, int) {
    sum := x + y
    difference := x - y
    return sum, difference
}

func main() {
    sayHello()

    result := add(5, 3)
    fmt.Println("5 + 3 =", result)

    sum, diff := calculate(10, 4)
    fmt.Println("Sum:", sum, "Difference:", diff)
}
