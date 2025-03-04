


package main

import "fmt"

func Add(a int, b int) int {
    return a + b
}

func Subtract(a int, b int) int {
    return a - b
}

func Multiply(a int, b int) int {
    return a * b
}

func Divide(a int, b int) int {
    return a / b
}

func main() {
    fmt.Println("What do you want to calculate:")

    var operation string
    fmt.Scan(&operation)

    if operation == "add" || operation == "Add" {
        fmt.Println("Please enter two numbers:")

        var a int
        var b int
        fmt.Scanf("%d", &a)
        fmt.Scanf("%d", &b)

        fmt.Println(Add(a, b))
    } else if operation == "subtract" || operation == "Subtract" {
        fmt.Println("Please enter two numbers:")

        var a int
        var b int
        fmt.Scanf("%d", &a)
        fmt.Scanf("%d", &b)

        fmt.Println(Subtract(a, b))
    } else if operation == "multiply" || operation == "Multiply" {
        fmt.Println("Please enter two numbers:")

        var a int
        var b int
        fmt.Scanf("%d", &a)
        fmt.Scanf("%d", &b)

        fmt.Println(Multiply(a, b))
    } else if operation == "divide" || operation == "Divide" {
        fmt.Println("Please enter two numbers:")

        var a int
        var b int
        fmt.Scanf("%d", &a)
        fmt.Scanf("%d", &b)

        fmt.Println(Divide(a, b))
    } else {
        fmt.Println("Invalid operation")
    }
}