package main

import "fmt"

func main() {
	// slices 엄청 많이 사용한다. 정확한 갯수를 모르니까.
	cars := []string{}
	cars = append(cars, "C Series")
	cars = append(cars, "E Series", "S Series", "M Series")
	cars[2] = "T Series"
	cars = append(cars, "X Series")

	fmt.Println(cars)
	// fmt.Println(names[2])
	// String default value는 nil이 아니라 "" empty string 이다.
	// fmt.Println("Cars[2] is nil: ", cars[2] == nil)
	fmt.Println("Cars[2] is nil: ", cars[2] == "")
}
