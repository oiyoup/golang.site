package main

import "fmt"

// 메서드를 전혀 갖지 않는 빈 인터페이스로서, Go의 모든 타입은 적어도 0개의 메서드를 구현하므로, 흔히
// Go에서 모든 타입을 나타내기 위해 빈 인터베이스를 사용한다.
// 빈 인터페이스는 어떠한 타입도 담을 수 있는 컨테이너라고 볼 수 있으며, 여러 다른 언어에서 흔히 일컫는
// 다이나믹 타입이라고 볼 수 있다.

func main() {
	var x interface{}
	// x = 1
	x = "Tom"

	printIt(x)

	var a interface{} = 1

	i := a
	j := a.(int)

	println(i)
	println(j)
}

func printIt(x interface{}) {
	fmt.Println(x)
}
