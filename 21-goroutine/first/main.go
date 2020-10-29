package main

import (
	"fmt"
	"time"
)

// Go 런타임은 Go루틴을 관리하면서 Go 채널을 통해 Go루틴 간의 통신을 쉽게 할 수 있도록 하였다.

func say(str string) {
	for i := 0; i < 10; i++ {
		fmt.Println(str, "***", i)
	}
}

func main() {
	say("Sync")

	go say("Async1")
	go say("Async2")
	go say("Async3")

	time.Sleep(time.Second * 3)
}
