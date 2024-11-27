package main

import "fmt"

func isNumberPar(number int) bool {
	return number%2 == 0
}

func isValidUser(nick string, password string) bool {
	return nick == "cris" && password == "12345"
}

func switch1(name string) {
	switch gods := name; gods {
	case "loki":
		fmt.Println("No thor")
	case "thor":
		fmt.Println("is thor")
	default:
		fmt.Println("gods...")
	}
}

func switch2(number int) {
	switch {
	case number > 100:
		fmt.Println("> 100")
	case number > 80 && number < 90:
		fmt.Println("> 90 < 80.")
	default:
		fmt.Println("Default")
	}
}

func main() {
	if isNumberPar(5) {
		fmt.Println("Es par")
	} else {
		fmt.Println("Es impar")
	}

	if isValidUser("cris", "12345") {
		fmt.Println("User ok")
	} else {
		fmt.Println("User not found")
	}

	switch1("thor")
	switch1("loki")
	switch1("hela")

	switch2(102)
	switch2(85)
	switch2(10)
}
