package game

import (
	"math/rand"
	"time"
)

type Deck struct {
	cards []Card
}

func CreateDeck() *Deck {
	deck := &Deck{}
	deck.cards = make([]Card, 52)
	var sindex Suit
	var nindex Face
	var cindex uint8 = 0
	for sindex = 1; sindex <= Hearts; sindex += 1 {
		for nindex = 1; nindex <= Ace; nindex += 1 {
			deck.cards[cindex] = Card{suit: sindex, face: nindex}
			cindex += 1
		}
	}
	return deck
}

func (d *Deck) GetCards() []Card {
	return d.cards
}

func (d *Deck) String() string {
	var result string = ""
	for index, c := range d.cards {
		result += c.ToSymbolString()
		if index+1 < len(d.cards) {
			result += ","
		}
	}
	return result
}

func (d *Deck) Shuffle(rounds int) {
	for round := 0; round < rounds; round += 1 {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(d.cards), func(i, j int) {
			d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
		})
	}
}

func (d *Deck) Draw(count int) Hand {
	cards := make([]Card, count)
	copy(cards, d.cards[:count])
	d.cards = d.cards[count:]
	return NewHand(cards)
}

func (d *Deck) Count() int {
	return len(d.cards)
}
