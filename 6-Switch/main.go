package main

import "fmt"

func main() {

	// if XXXXX {
	// 	WORK
	// } else if yyyyy {
	// 	WORK2
	// } else {
	// 	WORK3
	// }

	// SWITCH

	// switch var {
	// case var == condition:
	// 	WORK1
	// case var == condition2:
	// 	WORK2
	// default:
	// 	WOR K3
	// }

	tempratures := []int{10, -20, 30, 40}

	if tempratures[0] > 0 && tempratures[0] < 10 {
		fmt.Println("yekam sard")
	} else if tempratures[0] == 12 {
		fmt.Println("damaye aaali")
	} else if tempratures[0] >= 10 && tempratures[0] < 20 {
		fmt.Println("khonak")
	}

	switch tempratures[0] {
	case 12:
		fmt.Println("damaye aaali")
	case 30:
		fmt.Println("garm vali ghabele tahamol")

	default:
	}

	day := 1
	switch day {
	case 1:
		fmt.Println("Saturday")
	case 2:
		fmt.Println("Sunday")
	case 3:
		fmt.Println("Monday")
	case 4:
		fmt.Println("Tuesday")
	case 5:
		fmt.Println("Wednesday")
	case 6:
		fmt.Println("Thursday")
	case 7:
		fmt.Println("Friday")
	default:
		fmt.Println("Wrong day number")
	}

}
