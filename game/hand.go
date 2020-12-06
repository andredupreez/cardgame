package game

type Hand struct {
	cards []Card
}

func NewHand(cards []Card) Hand {
	return Hand{cards: cards}
}

func (h Hand) GetCards() []Card {
	return h.cards
}

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
