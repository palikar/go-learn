package main

import "fmt"

func main() {
	num := 20
OUTER:
	for true {
		num++
		for i := 1; i < 20; i++ {
			if num%i != 0 {
				continue OUTER
			}
		}
		fmt.Printf("%d\n", num)
		return

	}

}
