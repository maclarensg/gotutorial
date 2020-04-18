package main

import "fmt"

func main() {

	for i := 0; i <= 4; i++ {
		if i == 0 {
			fmt.Printf("i          i*2        i*3\n")
		}
		if i == 1 {
			fmt.Printf("==         ===        ===\n")
		}
		if i > 1 {
			fmt.Printf("%d          %d          %d\n", i, i*2, i*3)
		}
	}
}
