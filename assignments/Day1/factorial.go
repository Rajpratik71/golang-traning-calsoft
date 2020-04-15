package main

import "fmt"

var n int
var factVal uint64 = 1

func main() {
	fmt.Print("Enter a positive integer : ")
	fmt.Scan(&n)
	fmt.Print("Factorial is: ", factorial(n))
}

func factorial(n int) uint64 {
	if n < 0 {
		fmt.Print("Factorial of negative number doesn't exist.")
	} else {
		for i := 1; i <= n; i++ {
			factVal *= uint64(i)
		}

	}
	return factVal
}
