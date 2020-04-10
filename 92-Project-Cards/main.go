package main

import (
	"fmt"
	"os"
)

func main() {
	var mydeck deck
	draw := 4

	// Load deck either from file or new
	if fileExists("mydeck.txt") {
		mydeck = newDeckfromFile("mydeck.txt")
	} else {
		mydeck = newDeck()
	}

	// when deck is less than drawable cards, reshuffle
	if len(mydeck) < draw {
		fmt.Println("Reshuffle")
		mydeck = newDeck()
		mydeck.shuffle()
	}

	hand, remaining := mydeck.deal(draw)
	hand.print()
	remaining.saveDeck("mydeck.txt")
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
