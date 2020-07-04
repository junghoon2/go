package main

import "fmt"

func main() {
	// names := [3]string{"Keith", "Kevin", "Kayla"}
	// names := [3]string
	var cars [3]string

	cars[0] = "Sonata"
	cars[1] = "Grandeur"
	// cars[2] = "Avante"

	fmt.Println(cars)
	// fmt.Println(names[2])
	// String default value는 nil이 아니라 "" empty string 이다.
	// fmt.Println("Cars[2] is nil: ", cars[2] == nil)
	fmt.Println("Cars[2] is nil: ", cars[2] == "")
}
