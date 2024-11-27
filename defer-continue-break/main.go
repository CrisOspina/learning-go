package main

import "fmt"

func continueAndBreakExample() {
	for i := 0; i < 10; i = i + 1 {
		fmt.Println(i)

		if i == 5 {
			fmt.Println("continue")
			continue
		}

		if i == 8 {
			fmt.Println("break")
			break
		}
	}
}

func deferExample() {
	defer fmt.Println("Hi - 1")
	fmt.Println("Hi - 2")
	/*
		console
		Hi - 2
		Hi - 1
	*/
}

func main() {
	defer deferExample()
	continueAndBreakExample()
}
