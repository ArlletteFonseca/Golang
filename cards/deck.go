package main

import ("fmt"
		"strings"
		"io/ioutil"
		"os"
	
)
//Create a new type of 'deck'
//which is slice of strings
type deck []string
 
//new deck function
func newDeck() deck {
	cards := deck{}
	//these are two lists aka slices in golang
	cardSuits := []string{"Spades", "Diamonds","Hearts","Clubs"}
	cardValues := []string{"Ace","Two","Three", "Four"}

	// set up two four loops
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+ suit)
		}
	}
	return cards
}
 //d deck here is a receiver it is of type deck and function name is print
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
//call function of type deck and of type int handSize and d is name of variable or the argument
//(deck, deck) tells go you want to return to values of type deck, essentially string

func deal( d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string{
	strings.Join([]string(d), "")
	

}