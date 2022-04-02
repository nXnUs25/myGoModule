package myGoModule

import "fmt"

func PackageName() string {
	return "myPackage"
}

func CallPrivate(fname, lname string) {
	fmt.Println(privatePackage(fname, lname))
}

func Version() {
	fmt.Println("Version: ", version)
}

func Print() {
	fmt.Println("Package Calling print(): ")
	print()
}
