package main

import "fmt"

func main() {

	// c := 2

	// name := "Mahmoud"

	// Naghse | Rahnama
	// == : mosavi
	// != : na mosavi
	// >  : samte chap bayad az samte rast bozorgtar bashad
	// >= : samte chap bayad bozorgtar ya mosaviye samte rast bashad
	// <  : samte rast bayad az samte chap bozorgtar bashad
	// <= : samte rast bayad bozorgtar ya mosaviye samte chap bashad

	a := 10
	b := 5
	if a == b {
		fmt.Printf(" a mosavi ba b ast. a = %d, b = %d \n", a, b)
	} else if a < b {
		fmt.Printf(" a kochektar az b ast. a = %d, b = %d \n", a, b)
	} else if a > b {
		fmt.Printf(" a bozorgtar az b ast. a = %d, b = %d \n", a, b)
	} else {
		fmt.Printf(" marde hesabi dige chizi namoonde ke check koni! a= %d, b = %d \n", a, b)
	}

	// --------

	if a == b {
		fmt.Printf(" a mosavi ba b ast. a = %d, b = %d \n", a, b)
	}

	if a < b {
		fmt.Printf(" a kochektar az b ast. a = %d, b = %d \n", a, b)
	}

	if a > b {
		fmt.Printf(" a bozorgtar az b ast. a = %d, b = %d \n", a, b)
	}

	if a != b {
		fmt.Printf(" a mosavi ba b nist. a = %d, b = %d \n", a, b)
	}

	// ------
	d := false

	if d == false {
		fmt.Printf("meghdare d %t ast", d)
	}

	if d {
		fmt.Println("meghdare d true ast")
	}

	if !d {
		fmt.Println("meghdare d false ast")
	}

	if !(a == b) {
		fmt.Printf(" a mosavi ba b ast. a = %d, b = %d \n", a, b)
	}

	if !(a > b) {
		fmt.Printf(" a bozargtar az b NIST. a = %d, b = %d \n", a, b)
	}

	if a < b {
		fmt.Printf(" a kochektar az b AST. a = %d, b = %d \n", a, b)
	}

	// male: false, female: true
	names := []string{"delniya", "Mahmoud", "Yasin", "Shahrzad"}
	gender := []bool{true, false, false, true}

	fmt.Printf("Jensiyate %s %t ast \n", names[0], gender[0])
	fmt.Printf("Jensiyate %s %t ast \n", names[1], gender[1])
	fmt.Printf("Jensiyate %s %t ast \n", names[2], gender[2])
	fmt.Printf("Jensiyate %s %t ast \n", names[3], gender[3])

	if gender[0] {
		fmt.Printf("Jensiyate %s ZAN ast \n", names[0])
	} else if !gender[0] {
		fmt.Printf("Jensiyate %s MARD ast \n", names[0])
	}

	if gender[1] {
		fmt.Printf("Jensiyate %s ZAN ast \n", names[1])
	} else if !gender[1] {
		fmt.Printf("Jensiyate %s MARD ast \n", names[1])
	}

	if gender[2] {
		fmt.Printf("Jensiyate %s ZAN ast \n", names[2])
	} else if !gender[2] {
		fmt.Printf("Jensiyate %s MARD ast \n", names[2])
	}

	if gender[3] {
		fmt.Printf("Jensiyate %s ZAN ast \n", names[3])
	} else if !gender[3] {
		fmt.Printf("Jensiyate %s MARD ast \n", names[3])
	}

}
