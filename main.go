package main

import (
	"log"

	"github.com/andredupreez/cardgame/poker"
)

func main() {
	g := poker.CreatePoker()
	log.Printf("Shuffling...")
	g.Shuffle()
	hand := g.Draw()
	log.Printf("Your hand: %v", hand)
	_, description := g.Match(hand)
	log.Printf("You have: %v", description)
}
