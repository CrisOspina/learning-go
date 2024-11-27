package main

import "fmt"

func forBasic() {
	for i := 0; i < 10; i = i + 1 {
		fmt.Println(i)
	}
}

func forVar() {
	var i int = 0
	for i = 0; i < 10; i = i + 1 {
		fmt.Println(i)
	}
}

func while() {
	counter := 0

	for counter < 10 {
		fmt.Println((counter))
		counter++
	}
}

func forEver() {
	counterForever := 0
	for {
		fmt.Println((counterForever))
		counterForever = counterForever + 1
	}
}

func forDecrement() {
	for counter := 10; counter > 0; counter-- {
		fmt.Println(counter)
	}
}

func forRange() {
	list := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	for i, par := range list {
		fmt.Printf("Position %d number par: %d \n", i, par)
	}
}

func main() {
	forRange()
}
