package main

import "fmt"

func main() {

	// name1 := "Mahmoud"
	// name2 := "Hamid"
	// name3 := "Shahrzad"
	// name4 := "Yasin"
	// name5 := "Delniya"

	var xName []string
	fmt.Printf("%+v \n", xName)

	xName = append(xName, "Mahmoud")
	fmt.Printf("%+v \n", xName)

	xName = append(xName, "Hamid", "Shahrzad")
	fmt.Printf("%+v \n", xName)

	yName := []string{
		"Javad", // 0
		"Moein", // 1
		"Ario",  // 2
		"Arman", // 3
	}
	fmt.Printf("%+v \n", yName)

	fmt.Printf("Name man %s ast \n", yName[0])
	fmt.Printf("Name man %s ast \n", yName[1])
	fmt.Printf("Name man %s ast \n", yName[2])
	fmt.Printf("Name man %s ast \n", yName[3])

	// panic: runtime error: index out of range [4] with length 4
	// fmt.Printf("Name man %s ast \n", yName[4])

	zName := []string{}
	zName = append(zName, "Omid")  // 0
	zName = append(zName, "Leila") // 1
	zName = append(zName, "Ali")   //2
	fmt.Printf("zName shamele in item ha mishavad: %+v\n", zName)

	zName[2] = "Mohammad"
	fmt.Printf("zName shamele in item ha mishavad: %+v\n", zName)

	numbers := []int{}
	numbers2 := []int{1, 2, 3, 4, 5}
	numbers3 := []int{
		1,
		2,
		3,
		4,
		5,
	}

	fmt.Printf("Numbers: %+v \n", numbers)
	fmt.Printf("Second item in numbers2: %+v \n", numbers2[1])
	fmt.Printf("Numbers3: %+v \n", numbers3)

}
