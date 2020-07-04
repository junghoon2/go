package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"06.func/message"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Your Greeting: ")
	phrase, _ := reader.ReadString('\n')
	phrase = strings.TrimSpace(phrase)

	fmt.Print("Your Name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	m := message.Greeting(name, phrase)
	err := ioutil.WriteFile("/etc/motd", []byte(m), 0644)

	if err != nil {
		fmt.Println("Unable to write /etc/motd")
		os.Exit(1)
	}
}
