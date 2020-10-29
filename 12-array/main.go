package main

func main() {
    var a [3]int
    
    a[0] = 12
    a[1] = 24
    a[2] = 35
    
    println(a[0])
    println()
    
    for i := 0; i < len(a); i++ {
        println(i, a[i])
    }
    println()
    
    for idx, num := range a {
        println(idx, num)
    }
    println()
    
    bs := []int{1,2,3,4,5}
    
    for _, b := range bs {
        println(b)
    } 
}
