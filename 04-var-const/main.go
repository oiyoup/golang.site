package main

import "fmt"

func main() {
	var a int
	var f float32 = 11.
	// var i, j, k int
	// var i, j, k int = 1, 2, 3
	const c int = 10
	const (
		Vista  = "Visa"
		Master = "MasterCard"
		Amex   = "American Express"
	)
	const (
		Apple = iota
		Grape
		Orange
	)

	fmt.Println(a)
	fmt.Printf("%.3f\n", f)
	fmt.Println(c)
	fmt.Println(Vista)
	fmt.Println(Orange)
}
