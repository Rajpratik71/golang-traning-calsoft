package main

import "fmt"

func main() {
	studentRecord := make(map[int]string)
	studentRecord[01] = "Pratik raj"
	studentRecord[02] = "Pratik kumar"

	for key, value := range studentRecord {
		fmt.Print(" \n Roll number is : ", key)
		fmt.Print(" Student name is : ", value)
	}
}
