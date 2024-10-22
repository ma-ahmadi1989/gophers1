package main

import "fmt"

func main() {

	// AND   &&
	// OR    ||

	dama := 7

	// -10 <  x  < 5
	if dama > 10 && dama < 5 && dama < 10 {
		fmt.Println("salam")
	}

	if dama > 10 || dama < 5 || dama < 20 {
		fmt.Println("Mahmoud")
	}

	if dama > 15 && (dama < 10 || dama > 12) {
		fmt.Println("Yasin")
	}

}
