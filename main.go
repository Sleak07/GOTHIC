package main

import "fmt"

// conversion of farehnit to celcius
func kelvinToCelsius(k float64) float64 {
	k = k - 273.15
	return k
}

// first go function

func main() {
	fmt.Println("Hello after 6 years in GO")

	var myName string = "Arun"
	fmt.Println("My name is", myName)

	var number int = 10
	fmt.Println(number)

	var myAge uint8 = 6
	fmt.Println(myAge)

	// functions in go

	kelvin := 294.0
	celsius := kelvinToCelsius(kelvin)
	fmt.Print(kelvin, "° K is ", celsius, "° C")
}
