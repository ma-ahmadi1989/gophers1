package main

import "fmt"

func main() {

	// for ... ; ... ; ... {
	// 		our code comes here
	// }

	// a := 2
	// a = a + 1
	// a++

	for i := 1; i <= 789; i++ {
		baghimande := i % 3
		if baghimande == 0 {
			// fmt.Printf("addade %d bar 3 bakhshpazir ast \n", i)
			fmt.Println(i)
		}
	}

	// ----

	a := 0

	for i := 1; i <= 10; i++ {
		a = a + i
	}

	fmt.Printf("a is %d", a)

}
