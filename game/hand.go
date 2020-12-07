package game

type Hand struct {
	cards []Card
}

// construct a new hand
func NewHand(cards []Card) Hand {
	return Hand{cards: cards}
}

// Get hand cards
func (h Hand) GetCards() []Card {
	return h.cards
}

// Hand string description
func (h Hand) String() string {
	var result string = ""
	for index, c := range h.cards {
		result += c.ToSymbolString()
		if index+1 < len(h.cards) {
			result += ","
		}
	}
	return result
}
