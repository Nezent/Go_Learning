package main

func main() {
	cards := newDeck()
	cards.writeIntoFile("cards.txt")
}
