package main

import "fmt"

// conversion of Kelvin to Celsius
func kelvinToCelsius(k float64) float64 {
	return k - 273.15
}

// main function
func main() {
	fmt.Println("Hello after 6 years in GO")

	var myName string = "Arun"
	fmt.Println("My name is", myName)

	var number int = 10
	fmt.Println(number)

	var myAge uint8 = 6
	fmt.Println(myAge)

	// converting Kelvin to Celsius
	kelvin := 294.0
	celsius := kelvinToCelsius(kelvin)
	fmt.Printf("%.2f° K is %.2f° C\n", kelvin, celsius)

	// conditional statements
	command := "go east"
	if command == "go east" {
		fmt.Println("You head further up the mountain.")
	} else if command == "go inside" {
		fmt.Println("You enter the cave where you live out the rest of your life.")
	} else {
		fmt.Println("Didn't quite get that.")
	}
	// to get the user input
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	fmt.Printf("You entered: %f\n", input)
}
