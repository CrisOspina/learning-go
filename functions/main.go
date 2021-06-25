package main

import "fmt"

func test() string {
	const data string = "Hello"
	return data
}

func test2(a int) int {
	return a
}

func test3(a int) (c, d int) {
	return a, c
}

func main()  {
	dataA, dataB := test3(10)
	fmt.Println(test())
	fmt.Println(test2(12))
	fmt.Println(dataA, dataB)
}