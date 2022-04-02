package myGoModule

import "fmt"

const version = "2.0.0"

func OtherFileInSamePackage() string {
	return "Its other file location for this file"
}

func privatePackage(fname, lname string) string {
	return fname + " " + lname + " :Hello from version: " + version
}

func print() {
	fmt.Println("function been added and we didn't ran `go install`")
}
