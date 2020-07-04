package main

import "fmt"

func main() {
	// names := [3]string{"Keith", "Kevin", "Kayla"}
	// names := [3]string
	// var cars [3]string

	// cars[0] = "Sonata"
	// cars[1] = "Grandeur"
	// // cars[2] = "Avante"

	// cars := []string{}
	// cars = append(cars, "C Series")
	// cars = append(cars, "E Series", "S Series", "M Series")
	// cars[2] = "T Series"
	// cars = append(cars, "X Series")

	cars := make([]string, 4)
	cars[0] = "Tusan"
	cars[1] = "Palisade"
	cars[2] = "Vera"

	fmt.Println(cars)
	// fmt.Println(names[2])
	// String default value는 nil이 아니라 "" empty string 이다.
	// fmt.Println("Cars[2] is nil: ", cars[2] == nil)
	fmt.Println("Cars[2] is nil: ", cars[2] == "")
}
