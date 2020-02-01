package main

import (
	"fmt"
)

func is_prime(num int) bool {

	for i := 2; i < num/2; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {

	var sum = 0
	for i := range [1000]int{} {

		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Println("Sum: ", sum)

}
