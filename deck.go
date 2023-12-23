package main

import "fmt"

type deck []string

func newDeck() deck {
	card := deck{}
	cardSuit := []string {"Spades","Diamonds","Hearts","Clubs"}
	cardValue := []string {"Ace","Two","Three","Four"}

	for _,suit := range cardSuit {
		for _,value := range cardValue {
			card = append(card, value+" of "+suit)
		}
	}
	return card
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}