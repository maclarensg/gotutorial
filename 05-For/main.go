package main

import "fmt"

func main() {

	i := 1

	fmt.Println("for i <=3:")
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	fmt.Println("")

	fmt.Println("j := 7; j <= 9; j++")
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	fmt.Println("")

	fmt.Println("for")
	for {
		fmt.Println("loop")
		break
	}
	fmt.Println("")

	fmt.Println("n := 0; n <= 5; n++, with continue when mod 2")
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
	fmt.Println("")
}
