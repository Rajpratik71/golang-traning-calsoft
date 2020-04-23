package main

import "fmt"

func main() {
	myTechnologySlice := []string{"Docker", "Linux", "Python", "Cloud"}

	for _, technologies := range myTechnologySlice {
		fmt.Print("\n My Techology are : ", technologies)
	}
}
