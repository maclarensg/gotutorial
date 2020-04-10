package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

type deck []card

// (deck) print (): prints  all card
func (d deck) print() {
	for i, card := range d {
		fmt.Printf("%d: %s\n", i+1, card.toStr())
	}
}

// newDeck(): create a new deck of card
func newDeck() deck {
	cards := deck{}

	values := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "K", "Q"}
	houses := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	for _, h := range houses {
		for _, v := range values {
			cards = append(cards, card{v, h})
		}
	}
	return cards
}

// (deck) deal: return the a handsize of draw cards and the remaining cards
func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// (deck) shuffle: return a deck of shuffled cards
func (d deck) shuffle() {
	numCards := len(d)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := 0; i <= numCards-1; i++ {
		tmp := d[i]
		j := r1.Intn(numCards)
		d[i] = d[j]
		d[j] = tmp
	}
}

// (deck) toStr: Convert deck info to string
func (d deck) toStr() string {
	total := ""
	for i, card := range d {
		if i > 0 {
			total = total + ","
		}
		total = total + card.value + ":" + card.house
	}
	return total
}

// (deck) saveDeck: save deck to file
func (d deck) saveDeck(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toStr()), 0666)
}

// newDeckfromFile: load deck from file
func newDeckfromFile(filename string) deck {

	cards := deck{}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	cardsData := strings.Split(string(data), ",")

	if len(cardsData) > 0 {
		for _, c := range cardsData {
			cdata := strings.Split(string(c), ":")
			if len(cdata) == 2 {
				cards = append(cards, card{cdata[0], cdata[1]})
			}
		}
	}
	return cards
}
