
package main

import "fmt"

func main() {
    amongus := "hello defer"

    defer func() {
        fmt.Println(amongus)
    }()

    fmt.Println("hello1")
    fmt.Println("hello2")
    fmt.Println("hello3")
}