package main

import "fmt"

func main() {
    //first()
    //second()
    //third()
    //forth()
    //fifth()
    sixth()
}

func sixth() {
    ch := make(chan int, 2)
    
    ch <- 23
    ch <- 12
    
    close(ch)
    
    println(<- ch)
    println(<- ch)
    
    if _, success := <-ch; !success {
        println("No more data")
    }
}

func fifth() {
    ch := make(chan string, 1)
    sendChan(ch)
    receiveChan(ch)
}

func receiveChan(ch <-chan string) {
    data := <- ch
    fmt.Println(data)
}

func sendChan(ch chan<- string) {
    ch <- "hello"
}

func forth() {
    c := make(chan int, 1)
    
    c <- 23
    
    fmt.Println(<- c)
}

func third() {
    c := make(chan int)
    //c <- 1
    go func() {
        c <- 1
    }()
    fmt.Println(<- c)
}

func second() {
    done := make(chan bool)
    
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(i)
        }
        done <- true
    }()

    if <- done {
        fmt.Println("Done")
    }
}

func first()  {
    ch := make(chan int)

    go func() {
        ch <- 23
    }()

    fmt.Println(<- ch)
}
