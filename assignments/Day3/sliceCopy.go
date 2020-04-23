package main

import "fmt"

func main() {
	myTechnologySlice := make([]string, 5)
	fmt.Print("\n Before updating My Technology are : ", myTechnologySlice)
	myTechnologySlice = append(myTechnologySlice, "Docker", "Linux")
	fmt.Print("\n My Technology are : ", myTechnologySlice)
	myNewTechnologySlice := append(myTechnologySlice, "Python", "Cloud")
	fmt.Print("\n My New Technology are : ", myNewTechnologySlice)
}
