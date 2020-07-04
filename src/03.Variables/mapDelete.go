package main

import "fmt"

func main() {
	birthdays := map[string]string{
		"Keith": "02/06/1990",
		"Kevin": "01/01/1957",
		"Kayla": "06/24/1975",
	}

	delete(birthdays, "Keith")

	fmt.Println(birthdays)

	ages := map[string]int{}
	ages["Keith"] = 28
	ages["Kevin"] = 61
	ages["Kayla"] = 43

	fmt.Println(ages["Kevin"])
}
