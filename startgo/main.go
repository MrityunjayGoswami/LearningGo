package main


//Every element in slice should be same data type

func main() {

	cards := deck{newCard(), "new year"}
	cards = append(cards, "is coming.")

	// for i, card := range cards{

	// 	fmt.Println(i,card)
	// }
	cardofdeck := newDeck()
	cards.Print()
	cardofdeck.Print()
	

}

func newCard() string {
	return "5 diamonds"
}
