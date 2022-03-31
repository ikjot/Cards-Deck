package main

func main() {

	// Go is statically typed language. Similar to Java, C++, Scala.
	// Python, Ruby, Javascript is dynamically typed language.

	// Variable assignment (Long form)
	// var card string = "Ace of Spades"
	// Another way (abbreviated) :
	// card := "Ace of Spades"
	// := is used only for new variable initialization but not reassignment

	// fmt.Println(card)

	// Reassignment:
	// card = "Five of Diamonds"
	// fmt.Println(card)

	// Basic types of variable: bool, string, int, float64 etc.

	// newCard := newCard()
	// fmt.Println(newCard)

	// To handle list of records:
	// 	1) Array: Fixed length
	// 	2) Slice: Can grow or shrink at will. they are zero-indexed.
	// Both must be defined with a type.

	// Slice.
	// cards := []string{card, newCard, newCard}
	// run it as "go run main.go deck.go"
	// cards := deck{card, newCard, newCard}
	// cards = append(cards, "Six of Spades")
	// fmt.Println(cards)

	// cards := newDeck()
	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()

	// Type conversion
	// []<type we want>(object to convert)
	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))
	// Output: [72 105 32 116 104 101 114 101 33]

	// deck -> []string -> string -> []byte -> call WriteFile func
	// cards := newDeck()
	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()

}

// Go is not OO
