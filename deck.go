package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of deck which is a slice of strings.
// This is a way of saying: we are declaring a new type "deck" which
// inherits/extends all the behaviour of type "slice of strings"
type deck []string

// function to print each card

// (d deck) -> receiver of a function.
// Any variable of type deck now gets an access to print method.
// d -> reference to instance of deck variable. It is similar to "this" and "self" in other languages.
// By convention: use 2-3 letters of variable
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suite := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
		}
	}
	return cards
}

// func (d deck) deal(handSize int).. -> did not work as we need pointer -> func (dPtr *Deck) deal(handSize int) Deck {
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	// if something went wrong: err will be populated
	// it will be nil, if everything went all good

	if err != nil {
		// Error occurred.
		// Option #1 - Log the error and return a call to newDeck().
		// Option #2 - Log the error and entirely quit the program.
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(bs), ","))
}

// func (d deck) shuffle() {
// 	l := len(d)
// 	for i, _ := range d {
// 		newPosition := rand.Intn(l - 1)
// 		d[i], d[newPosition] = d[newPosition], d[i]
// 	}
// }

// This function always prints the same random order:
// $ go run main.go deck.go
// 0 Two of Hearts
// 1 Ace of Spades
// 2 Three of Spades
// 3 Two of Diamonds
// 4 Ace of Clubs
// 5 Four of Diamonds
// 6 Three of Hearts
// 7 Three of Clubs
// 8 Ace of Diamonds
// 9 Four of Hearts
// 10 Four of Spades
// 11 Ace of Hearts
// 12 Two of Spades

// something is not quite right. rand.Intn -> pseudo random generator.
// newPosition := rand.Intn(l - 1)
// It uses a seed and by default it uses the same exact seed value.
// Use a random seed.

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	l := len(d)
	for i, _ := range d {
		newPosition := r.Intn(l - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
