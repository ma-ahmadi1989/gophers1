package main

import "fmt"

func main() {
	var object string
	var gender bool

	fmt.Println("==[ Objects ]==")

	object = "Window"
	fmt.Println("1) ", object)

	object = "chair"
	fmt.Println("2) ", object)

	fmt.Println("")
	fmt.Println("")
	fmt.Println("==[ Genders (Male: false, Female: true) ]==")

	gender = false
	fmt.Println("1) ", gender)

	gender = true
	fmt.Println("2) ", gender)

}
