package main

//import "fmt"

func main() {
	cards := newDeck()
	cards.saveToFile("saveTest")

	//hand, remainingCards := deal(cards, 5)

	//hand.print()
	//fmt.Println("------------")
	//remainingCards.print()
	//remainingCards.saveToFile()
	//fmt.Println(hand.toString())
}


