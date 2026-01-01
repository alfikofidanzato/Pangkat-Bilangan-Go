package main

import (
	"fmt"
	"time"
)

func powerIterative(x int, n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result *= x
	}
	return result
}

func powerRecursive(x int, n int) int {
	if n == 0 {
		return 1
	}
	temp := powerRecursive(x, n/2)
	if n%2 == 0 {
		return temp * temp
	} else {
		return x * temp * temp
	}
}

func main() {
	x := 10 

	fmt.Println("x =", x)
	fmt.Println("n\tIteratif(us)\tRekursif(us)")

	for _, n := range []int{100, 500, 1000, 5000, 10000, 20000, 50000, 100000} {

		startIter := time.Now()
		powerIterative(x, n)
		elapsedIter := time.Since(startIter)

		startRec := time.Now()
		powerRecursive(x, n)
		elapsedRec := time.Since(startRec)

		fmt.Printf("%d\t%d\t\t%d\n", n, elapsedIter.Microseconds(), elapsedRec.Microseconds())
	}
}
