package main

import "fmt"
//go is not object oriented with classes 
func main(){
	// var card string = "Ace of Spades"
	// card:= "Ace of Spades" //colon initializes value
	// cards := deck{"Ace of Diamonds", newCard()}
	//adding a new card inside
	// cards = append(cards, "Six of Spades")
	//golang is static type
	//iterate slice
	// for i, card := range cards {
	// 	fmt.Println(i,card)
	// }
	// cards := newDeck()
	//here I am assigning hand to cards and remaining cards to 5
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()
	cards := newDeck()
	fmt.Println(cards.toString())
}
