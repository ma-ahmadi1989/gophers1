package main

import "fmt"

func main() {
	objectList := []string{"pillow", "chair", "lamp"}
	damaha := []int{12, 15, 32, -5, 17, 19}

	fmt.Printf("list ashya: %v \n", objectList)

	fmt.Printf("damaye 1: %d, damaye 2: %d \n", damaha[0], damaha[1])
	fmt.Printf("damaye 3: %d, damaye 4: %d \n", damaha[2], damaha[3])

	// male: false, female: true
	names := []string{"Mahmoud", "Yasin", "Delniya", "Shahrzad"}
	gender := []bool{false, false, true, true}

	fmt.Printf("jensiyate %s %t ast \n", names[0], gender[0])
	fmt.Printf("jensiyate %s %t ast \n", names[1], gender[1])
	fmt.Printf("jensiyate %s %t ast \n", names[2], gender[2])
	fmt.Printf("jensiyate %s %t ast \n", names[3], gender[3])

}
