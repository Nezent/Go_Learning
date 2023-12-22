package main

import "fmt"

func main() {
	cards := []string{}
	cards = append(cards, "Aces of Hearts")
	cards = append(cards, "Ace of Diamonds")
	cards = append(cards, "Six of Spades","Six of Hearts")
	for i, card := range cards {
		fmt.Println(i,card);
	}
}

