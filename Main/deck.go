package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// The parenthesis before the function is called a "reciever". It is essentially
// a way to add methods to something. In this case we are adding a print method
// to type deck. explanation ("this value" type of)
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// This reciever takes a slice of strings and turns it into a single string
// using the strings.join method from the standard library.
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Save the file and give the program read & write permission. (0666)
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// Randomize the deck of cards.
func (d deck) shuffle() {
	/*
		1.) Use time method to get a new int64 number and create the seed.
		2.) Use rand.NewSource to generate a new source.
		3.) Take the source to generate a new randomizer.
	*/
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
