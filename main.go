package main

import (
	"fmt"

	"github.com/andredupreez/cardgame/poker"
)

func main() {
	g := poker.CreatePoker()
	fmt.Printf("Shuffling...\n")
	g.Shuffle()
	hand := g.Draw()
	fmt.Printf("Your hand: %v\n", hand)
	_, description := g.Match(hand)
	fmt.Printf("You have:  %v\n", description)
}
