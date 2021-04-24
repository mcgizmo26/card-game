package main

import "fmt"

func main() {
	cards := newDeckFromFile("my_deck")
	cards.shuffle()

	hand1, hand2 := deal(cards, 5)

	fmt.Println("My 1st hand is " + hand1.toString())
	fmt.Println("My 2nd hand is " + hand2.toString())
}
