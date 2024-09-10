package main

import "fmt"

func main2() {
	adade1 := 100
	fmt.Println(adade1)

	var adade2 uint64
	adade2 = 18446744073709551614
	fmt.Println(adade2)

	adade2 = 1000
	fmt.Println(adade2)

	var adade3 uint64 = 18446744073709551614
	fmt.Println(adade3)

	var name string
	name = "Mahmoud"
	fmt.Println(name)

	var check1 bool
	check1 = false
	fmt.Println(check1)

	var check2 bool = true
	fmt.Println(check2)

	check3 := false
	fmt.Println(check3)
}
