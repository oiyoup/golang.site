package main

import (
	"fmt"

	"./first"
	"./second"
	// _ "./first" // init만 호출할 때
)

func init() {
	fmt.Println("Init")
}

func main() {
	fmt.Println("Main")
	first.Good()
	first.Handler()
	second.Good()
}
