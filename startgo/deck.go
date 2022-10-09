package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"spades","diamonds","hearts","clubs"}
	cardValues := []string{"Ace","2","3","4","5"}
	
	for _ , suit := range cardSuits{
		for _, value := range cardValues{
			cards = append(cards, value +" of "+suit)
		}
	}
	
	return cards
}

func (d deck) Print(){

	for i, card := range d{

		fmt.Println(i,card)
	}
}
