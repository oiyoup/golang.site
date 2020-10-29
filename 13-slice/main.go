package main

import "fmt"

func main() {
    var a []int
    
    a = []int{2, 5, 9}
    a[2] = 23
    fmt.Println(a)
    
    s := make([]int, 5, 10)
    fmt.Println(s)
    fmt.Println(len(s), cap(s))
    
    s = []int{12, 34, 45, 56, 67}
    fmt.Println(s)
    fmt.Println(s[:2])
    fmt.Println(s[2:])
    fmt.Println(s[:])
    
    s = append(s, 99)
    if s == nil {
        fmt.Println("Same")
    }
    
    s1 := []int{1, 2, 3}
    
    s1 = append(s, s1...)
    fmt.Println(s1)
    
    source := []int{0, 1, 2}
    target := make([]int, len(source), cap(source)*2)
    copy(source, target)
    fmt.Println(target)
}
