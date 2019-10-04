package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Create a new type of 'deck' slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"♥", "♣", "♠", "♦"}
	cardValues := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "V", "D", "K"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" "+suit)
		}
	}

	cards = append(cards, "Joker")

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) printAll() {
	fmt.Println(d)
}

func (d deck) deal(size int) (deck, deck) {
	return dealCards(d, size)
}

func (d deck) convertToBytes() []byte {
	return []byte(d.toString())
}

func (d deck) toString() string {
	// return strings.Join([]string(d), ",")
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, d.convertToBytes(), 0666)
}

// --- helpers
func dealCards(d deck, size int) (deck, deck) {
	return d[:size], d[size:]
}
