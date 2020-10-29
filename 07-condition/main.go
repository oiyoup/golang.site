package main

import "fmt"

func main() {
    //k := 12
    
    //If1(k)
    //if2(k)
}

func if2(k int) {
    max := 45
    if val := k * 2; val < max {
        fmt.Println("hello")
    }
    
    //println(val)
}

// If1 is good
func If1(k int) {
    if k == 1 {
        fmt.Println("One")
    } else if k == 2 {
        fmt.Println("Two")
    } else {
        fmt.Println("Other")
    }
}