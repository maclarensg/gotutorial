package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	var c float32
	prompt := color.CyanString("Celsius (°C): ")
	fmt.Printf(prompt)

	// Get Celsius input
	_, err := fmt.Scanf("%f", &c)

	if err != nil {
		fmt.Println("Invalid Response that is not of type Float")
	}
	color.Blue("Fahrenheit (°F): %f", (c/5*9)+32)
}
