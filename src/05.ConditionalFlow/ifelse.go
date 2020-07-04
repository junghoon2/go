package main

import "fmt"

func main() {
	ages := map[string]int{}
	ages["Kevin"] = 50

	// if 문에서는 조건이 한번만 실행된다.
	if ages["Kevin"] < 20 {
		fmt.Println("Kevin can vote, exit 1")
	} else if ages["Kevins"] < 65 { // 앞의 조건이 실행되면 if 밖으로 나가 버린다.
		fmt.Println("Work Hard")
	} else {
		fmt.Println("Enjoy your retirement age")
	}
}
