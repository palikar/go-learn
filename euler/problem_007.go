package main

import "fmt"

func is_prime(num int) bool {

	for i := 2; i < num/2; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	prime := 2
	ind := 3
	cnt := 0
	for cnt < 10001 {
		if is_prime(ind) {
			prime = ind
			cnt++
		}
		ind++
	}
	fmt.Printf("%d\n", prime)

}
