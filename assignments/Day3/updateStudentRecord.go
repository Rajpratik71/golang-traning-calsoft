package main

import "fmt"

func main() {
	studentRecord := make(map[int]string)
	studentRecord[01] = "Pratik raj"
	studentRecord[02] = "Pratik kumar"

	fmt.Print(" Before modifying Entries : ")

	for key, value := range studentRecord {
		fmt.Print(" \n Roll number is : ", key)
		fmt.Print(" Student name is : ", value)
	}

	studentRecord[01] = "Pratik kumar"
	studentRecord[02] = "Pratik raj"

	fmt.Print("\n After modifying Entries : ")

	for key, value := range studentRecord {
		fmt.Print(" \n Roll number is : ", key)
		fmt.Print(" Student name is : ", value)
	}

}
