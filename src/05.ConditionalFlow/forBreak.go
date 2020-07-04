package main

import "fmt"

func main() {
	// initial 먼저 하고 다음 조건문
	// for 수행하고 해야 할 것
	// i for iterable
	a := 0
	for a < 10 {
		if a%2 == 0 {
			a++
			continue // continue for next iteration
		} else if a == 5 {
			break // for 문으로 완전 종료
		}
		fmt.Println("We're counting(again): ", a)
		// for 문 안에 incremental 삽입 가능
		a++
	}
}
