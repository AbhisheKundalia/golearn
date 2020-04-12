package main

func main() {
	// card := "Ace of Spades"
	// card = newCard()
	cards := deck{"Ace of diamonds", newCard()}
	cards = append(cards, "Six of Spades")
	//fmt.Println(cards)
	cards.print()
}

func newCard() string {
	return "Five of diamonds"
}
