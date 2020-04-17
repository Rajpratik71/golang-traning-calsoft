package main

import (
	"encoding/json"
	"io/ioutil"
)

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

	writefile, _ := json.MarshalIndent(studentInformation, "", " ")

	_ = ioutil.WriteFile("test.json", writefile, 0644)

}
