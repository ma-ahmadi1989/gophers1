package main

import "fmt"

func main() {

	adade1 := 12

	adade2 := 19.5

	gender := false

	// fmt.Println
	fmt.Println("Hello World")
	fmt.Println("Hello", "World")

	fmt.Println(adade1)
	fmt.Println("adade1:", adade1)
	fmt.Println("adade1:", adade1+1)

	// fmt.Printf
	// %s => string
	// %d => digit
	// %f => float64
	// %t => bool
	// \n => new line
	fmt.Println("adade avvale man 1 ast va agar yek vahed be an ezafe konim 2 mishavad")
	fmt.Println("adade avvale man", 1, "ast va agar yek vahed be an ezafe konim", 2, "mishavad")
	fmt.Println("adade avvale man", adade1, "ast va agar yek vahed be an ezafe konim", adade1+1, "mishavad")
	fmt.Printf("adade avvale man %d ast va agar yek vahed be an ezafe konim %d mishavad\n", adade1, adade1+1)
	fmt.Printf("adade 2 %f ast va gensiyat %t aster\n", adade2, gender)

}
