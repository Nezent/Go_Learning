package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

func deal(d deck,handSize int) (deck,deck) {
	return d[:handSize],d[handSize:]
}

func (d deck) writeIntoFile(filename string) {

	// Open the file for writing, create it if it doesn't exist
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Create a bufio.Writer for efficient writing
	writer := bufio.NewWriter(file)

	// Write the string to the file
	_, err = writer.WriteString(d.toString())
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	// Flush the bufio.Writer to ensure all data is written to the file
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}

	fmt.Println("String has been successfully written to", filename)
}

func (d deck) toString() string {
	return strings.Join([]string(d),",")
}