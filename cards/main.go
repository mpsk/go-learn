package main

func main() {
	cards := newDeck()
	_, remainigCards := cards.deal(3)

	// hand.print()
	// remainigCards.print()

	remainigCards.saveToFile("file1cards")

	// fmt.Println(remainigCards.convertToBytes())
}
