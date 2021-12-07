package main

import "fmt"

func main() {
	//var card = newCard()
	card, i := newCard()
	fmt.Println(card, i)

}

func newCard() (string, int) {
	return "five", 3
}
