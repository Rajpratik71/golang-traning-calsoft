package main

import "fmt"

func main() {
	studentRecord := make(map[int]string)
	studentRecord[01] = "Pratik raj"
	studentRecord[02] = "Pratik kumar"

	fmt.Println(" Before modifying Entries : ")

	for key, value := range studentRecord {
		fmt.Print(" \n Roll number is : ", key)
		fmt.Print(" Student name is : ", value)
	}

	fmt.Println("\n After deleting Entries : ")

	for key, value := range studentRecord {
		fmt.Print(" \n Roll number is : ", key)
		fmt.Print(" Student name is : ", value)
	}

	defer fmt.Println("\n Earlier not able to delete Entries because defer is used. Now Student Entries are : ", studentRecord)

	defer delete(studentRecord, 01)

}
