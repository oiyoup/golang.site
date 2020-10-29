package main

import "fmt"

func main() {
    var i interface{} = 23
    
    j, ok := i.(int)
    if !ok {
        fmt.Println("Panic")
    }
    
    fmt.Println(j)
}
