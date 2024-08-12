package main

import "fmt"

func main() {

	// if and else
	// if a < 8 print: less than 8
	// if 8 < a > 10 print: between 8 and 10
	// if 10 < a > 12 print: betwrrn 10 and 12
	// if 12 < a print: more than 12
	// &: shift + 7 (Keyboard)
	// &&: AND
	// ||: OR
	// == : equal - mosavi
	// >= : bozorgtar-mosavi
	// <= : kochecktar-mosavi

	a := 11

	if a < 8 {
		fmt.Println("less than 8")
	}

	//
	if a > 8 && a < 10 {
		fmt.Println("between 8 and 10")
	}

	if a > 10 && a < 12 {
		fmt.Println("between 10 and 12")
	}

	if a > 12 {
		fmt.Println("more than 12")
	}

	if a == 10 {
		fmt.Println("a is 10")
	}

	if a > 10 || a < 20 {
		fmt.Println("greater than 10 or less than 20")
	}

	if a < 5 {
		fmt.Println("greater than 5")
	} else if a > 5 && a < 12 {
		fmt.Println("between 5 and 12")
	} else {
		fmt.Println("my name is javad")
	}

	a = 2 + 5
	a = 5 - 4
	a = 3 * 10
	a = 9 / 3
	a = 9 % 3

}
