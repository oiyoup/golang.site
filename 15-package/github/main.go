package main

import (
	"fmt"

	"github.com/oiyoup/package/first"
	"github.com/oiyoup/package/second"
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
