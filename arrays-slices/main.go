package main

import (
	"fmt"
	"strings"
)

// Los arrays poseen un tamaño fijo y son inmutables, mientras que en los slices su tamaño es dinámico y los puedes modificar.

func arrayExample() {
	var data [4]int
	data[0] = 1
	data[1] = 2
	data[2] = 3

	fmt.Println("total: ", len(data), "capacidad del array", cap(data))
	fmt.Println(data)
}

func slideExample() {
	slice := []int{0,1,2,3,4,5,6,7}
	fmt.Println("total: ", len(slice), "capacidad del slice: ", cap(slice))
}

func methodForSlice() {
	slice := []int{0,1,2,3,4,5,6,7}
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	slice = append(slice, 8)
	fmt.Println(slice)

	newSlice := []int{9, 10, 11}
	slice = append(slice, newSlice...)
	fmt.Println(slice)
}	

func sliceExample2() {
	data := []string{"thor", "loki", "hela"}

	for _, value := range data {
		fmt.Println(value)
	}
}

func isPalindromo1(text string) bool {
	var textReverse string

	for i := len(text) - 1; i >= 0; i-- {
		textReverse = textReverse + string(text[i])
	}

	return strings.ToLower(text) == strings.ToLower(textReverse)
}

func main()  {	
	arrayExample()
	slideExample()
	methodForSlice()
	sliceExample2()
	
	if(isPalindromo1("ama")) {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No lo es palindromo...")
	}
}