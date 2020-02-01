package main

import (
	"fmt"
)

// one_digit ...
func one_digit(i int) bool {
	return i >= 0 && i < 10
}

// is_pali ...
func is_pali(num int, dupNum *int) bool {

	if one_digit(num) {
		return num == *dupNum%10
	}

	if !is_pali(num/10, dupNum) {
		return false
	}

	*dupNum /= 10

	return num%10 == *dupNum%10
}

// pali ...
func pali(num int) bool {

	if num < 0 {
		num = -num
	}

	var dup int = num

	return is_pali(num, &dup)
}

func main() {

	max := 0
	for i := 100; i <= 999; i++ {
		for j := 100; j <= 999; j++ {
			prod := i * j

			if !pali(prod) {
				continue
			}

			if prod > max {
				max = prod
			}

		}
	}

	fmt.Printf("%d\n", max)

}
