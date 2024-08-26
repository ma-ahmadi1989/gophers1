package main

import "fmt"

func main() {

	// fibonacci
	first := 0
	second := 1

	fmt.Printf("fibonacci: %d,%d", first, second)

	var sum int

	for i := 0; i <= 100; i++ {
		sum = first + second
		fmt.Printf(",%d", sum)

		first = second
		second = sum

	}
	fmt.Println("")
	fmt.Println("")

	// zarb

	var aval int64
	var natijeh int64
	aval = 2

	fmt.Printf("Zarb: %d,", aval)

	for i := 1; i <= 39; i++ {
		natijeh = aval * 2
		fmt.Printf("%d,", natijeh)
		aval = natijeh
	}
	fmt.Println("")

}
