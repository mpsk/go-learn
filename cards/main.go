package main

func main() {
	filename := "file_cards";
	cards := newDeck()
	_, remainigCards := cards.deal(3)

	remainigCards.saveToFile(filename)

	newDeckFromFile(filename).print()
}
