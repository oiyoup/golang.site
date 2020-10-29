package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(2)

	var wait sync.WaitGroup
	wait.Add(2)

	go func() {
		defer wait.Done()
		fmt.Println("First")
	}()

	go func(msg string) {
		defer wait.Done()
		fmt.Println(msg)
	}("Second")

	wait.Wait()
}
