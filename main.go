package main

func getCard() string {
	return "My Very Own Sun!"
}

func main() {

	cards := createDeck()

	// taken, remain := cards.deal(5)

	// taken.printDeck()
	// fmt.Println("What remains?")
	// remain.printDeck()
	// fmt.Println(cards.toString())
	// cards := newDeckFromFile("my_deck")
	cards.printDeck()

	cards.shuffle()

	cards.printDeck()
}
