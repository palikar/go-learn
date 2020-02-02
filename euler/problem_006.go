package main

import "fmt"
import "math"

func main() {

	sum1 := 0.0
	sum2 := 0.0
	for i := 1; i <= 100; i++ {
		sum1 += math.Pow(float64(i), 2)
		sum2 += float64(i)
	}

	sum2 = math.Pow(sum2, 2)

	diff := sum2 - sum1
	fmt.Printf("%f\n", diff)

}
