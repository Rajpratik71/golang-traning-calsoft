package main

import "fmt"

type StudentStruct struct {
	firstName string
	lastName  string
	age       int8
}

func main() {

	studentInformation := StudentStruct{
		firstName: "Pratik",
		lastName:  "raj",
		age:       22,
	}

	fmt.Println(" Before Updaating Student Name is : ", studentInformation.firstName+" "+studentInformation.lastName, " and Student age is : ", studentInformation.age)

	studentInformation.updateStudentInformation()

	fmt.Print(" After Updaating Student Name is : ", studentInformation.firstName+" "+studentInformation.lastName, " and Student age is : ", studentInformation.age)
}

func (studentInformation *StudentStruct) updateStudentInformation() {

	*studentInformation = StudentStruct{
		firstName: "Pratik",
		lastName:  "raj",
		age:       23,
	}

	fmt.Println("Updating Student Information")
}
