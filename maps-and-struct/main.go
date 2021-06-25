package main

import "fmt"

func hashMapsExample() {
	data := make(map[string]int)
	data["thor"] = 100
	data["odin"] = 500
	data["hela"] = 200
	fmt.Println(data)
	
	for key, value := range data {
		fmt.Println(key, value)
	}
}

type car struct {
	brand string
	year int
}

func main() {
	hashMapsExample()
	myCar := car{brand: "Ford", year: 2000}
	fmt.Println(myCar)

	var otherCar car
	otherCar.brand = "Ferrari"
	otherCar.year = 1977
	fmt.Println(otherCar)
}