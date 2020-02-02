package main

import "fmt"

func pow10(num int) int {
	res := 1
	for i := 0; i < num; i++ {
		res *= num
		res = res % 10000000000
	}
	return res
}

func main() {

	res := 0
	for i := 1; i <= 1000; i++ {
		last := pow10(i)
		res += last
		res = res % 10000000000
	}

	fmt.Printf("%d\n", res)

}
