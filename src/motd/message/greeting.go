package message

import "fmt"

// Greeting hello
// import 되는 함수는 대문자로 시작한다.
func Greeting(name, message string) string {
	return fmt.Sprintf("%s, %s", message, name)
}
