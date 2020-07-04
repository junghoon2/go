package main

import (
	"fmt"

	"06.func/message"
)

func main() {
	m := message.Greeting("Keith", "Hello")
	fmt.Println(m)
}

// main 함수는 input, output이 없다. void function.
// func main() {
// 	m := message.greeting("Keith", "Hello")
// 	fmt.Println(m)
// 	// 변수를 함수 return 값으로 할 수 있다.
// 	// message := greeting("Keith", "Hello")
// 	// fmt.Println(message)
// }

// // 함수 선언하기
// // return 값을 가진다.
// // func keyword
// // func greeting(name string, message string) string {
// // 같은 type이면 string 2번 안 적어줘도 된다.
// func greeting(name, message string) (salutation string) {
// 	// Sprint가 아니라 Sprintf 이다.
// 	salutation = fmt.Sprintf("%s, %s", message, name)
// 	// return fmt.Sprintf("%s, %s", message, name)
// 	return
// }
