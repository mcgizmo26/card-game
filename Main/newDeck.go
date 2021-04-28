package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Parameters 1st value 2nd type then comma. Go convention is that you use
// a single letter to represent the "this" value. The 2nd set of parenthesis
// is the return type. Go can return more than one value.
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// create a deck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Daimonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}

	// Use _ (ie. underscore) whenever you need to have a variable but
	// you don't need to use it.
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Create a new deck from file.
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	// If there is an error print it and exit the program.
	if err != nil {
		tempDeck := newDeck()
		err2 := tempDeck.saveToFile("My_deck")
		if err2 != nil {
			fmt.Println("Error: ", err2)
			os.Exit(1)
		} else {
			nd := newDeckFromFile(filename)
			return nd
		}
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}
