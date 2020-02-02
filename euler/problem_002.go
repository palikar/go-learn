package main

import (
	"fmt"
	"math"
)

type Person struct {
	name string
	full string
}

func main() {

	var fib1 = 1
	var fib2 = 1

	var sum = 0

	for fib2 < 4000000 {

		if fib2%2 == 0 {
			sum += fib2
		}

		old2 := fib2
		fib2 = fib1 + fib2
		fib1 = old2

	}

	p := Person{name: "name", full: "name full"}

	fmt.Println("Sum:", sum)

}
