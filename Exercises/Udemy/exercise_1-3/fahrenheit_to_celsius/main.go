package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	var f float32
	prompt := color.CyanString("Fahrenheit (°F): ")
	fmt.Printf(prompt)

	// Get Fahrenheit input
	_, err := fmt.Scanf("%f", &f)

	if err != nil {
		fmt.Println("Invalid Response that is not of type Float")
	}
	color.Blue("Celsius (°C): %f", (f-32)*5/9)
}
