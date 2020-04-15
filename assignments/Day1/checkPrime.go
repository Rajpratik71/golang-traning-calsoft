package main

import "fmt"

var n, i int
var flag = 0

func checkPrime(n int) {

	// for i := 2; i <= n/2; i++ {

	// 	if n%i == 0 {
	// 		flag := 1
	// 		break
	// 	}

	// }

	if n == 1 {
		fmt.Print("1 is neither prime nor composite.")
	} else {
		if flag == 0 {
			fmt.Print("%d is a prime number.", n)
		} else {
			fmt.Print("%d is not a prime number.", n)
		}
	}

}

func main() {
	fmt.Print("Enter a Number :")
	fmt.Scan(&n)
	fmt.Print("The Number is :", checkPrime(n))

}
