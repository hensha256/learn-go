package main

import "fmt"

func Hello() string {
    return "Hello, World"
}

const helloPrefix = "Hello, "

func HelloYou(name string) string {
    if name == "" {
        name = "World"
    }
    return helloPrefix + name
}

func main() {
    fmt.Println(HelloYou("Alice"))
}
