package main

import (
	"fmt"
)

func main() {

	var fib_1 = 1
	var fib_2 = 1

	var sum = 0

	for fib_2 < 4000000 {

		if fib_2%2 == 0 {
			sum += fib_2
		}

		old_2 := fib_2
		fib_2 = fib_1 + fib_2
		fib_1 = old_2

	}

	fmt.Println("Sum:", sum)

}
