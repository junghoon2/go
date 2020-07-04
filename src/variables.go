package main

import "fmt"

func main() {
	// var myInt int = 16
	// var val, ok = "yes", true // multiple variable at same time
	var val, _ = "yes", true // not defined var name

	myInt := 8 // Same for declare variable
	name2 := "Paul"

	var name string // 먼저 type만 선언도 가능
	name = "Keith"

	fmt.Println("myInt is:", myInt)
	fmt.Println("myInt times two:", myInt*2)
	fmt.Println("val is:", val)
	//fmt.Println("ok is:", ok)
	fmt.Println("Variable name is:", name)
	fmt.Println("Another name is:", name2)
}
