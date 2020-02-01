package main

import (
	"fmt"
	"math"
)

func main() {
	var facts []int
	var num = 600851475143

	if num%2 == 0 {
		facts = append(facts, 2)
		num /= 2

	}

	for i := 3; i <= int(math.Sqrt(float64(num))); i += 2 {
		if num%i == 0 {
			facts = append(facts, i)
			num /= i
		}
	}

	if num > 2 {
		facts = append(facts, num)
	}

	fmt.Println(facts[len(facts)-1])

}
