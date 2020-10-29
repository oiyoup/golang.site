package main

func nextVal() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	next := nextVal()

	println(next())
	println(next())
	println(next())
	println()

	anotherNext := nextVal()

	println(anotherNext())
	println(anotherNext())
	println(anotherNext())
}
