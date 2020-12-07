package game

type Deck struct {
	cards []Card
}

// Create a new deck
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

// Get the list of cards in the deck
func (d Deck) GetCards() []Card {
	return d.cards
}

// Format the string representation of the deck and cards
func (d Deck) String() string {
	var result string = ""
	for index, c := range d.cards {
		result += c.ToSymbolString()
		if index+1 < len(d.cards) {
			result += ","
		}
	}
	return result
}

// Shuffle the deck for the specified number of
// rounds using the shuffle function specified
func (d *Deck) Shuffle(rounds int, fun func([]Card) []Card) {
	for round := 0; round < rounds; round += 1 {
		d.cards = fun(d.cards)
	}
}

// Draw the specified number of cards from the deck returned as a Hand,
// the deck is modified as the cards drawn is removed from the deck
func (d *Deck) Draw(count int) Hand {
	cards := make([]Card, count)
	copy(cards, d.cards[:count])
	d.cards = d.cards[count:]
	return NewHand(cards)
}

// Return the number of cards remaining in the deck
func (d Deck) Count() int {
	return len(d.cards)
}
