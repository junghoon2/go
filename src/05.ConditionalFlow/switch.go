package main

import "fmt"

func main() {
	ages := map[string]int{}
	ages["Kevin"] = 20

	// switch {
	// case ages["Kevin"] < 18:
	// 	fmt.Println("Kevin can't vote")
	// case ages["Kevin"] < 65:
	// 	fmt.Println("It is not of retirement age")
	// default:
	// 	fmt.Println("Enjoy your retirement age")
	// }

	switch ages["Kevin"] {
	case 1, 2, 3, 5, 7, 11, 13, 17, 19:
		fmt.Println("Kevin's age is a prime number!")
	// case ages["Kevin"] > 16:, 이렇게는 사용 못하는구나
	// 	fmt.Println("Kevin can drive now")
	case 18:
		fmt.Println("Kevin can vote now!")
	case 67:
		fmt.Println("Kevin can retire now!")
	default:
		fmt.Println("There's nothing special about Kevin's current age.")
	}
}
