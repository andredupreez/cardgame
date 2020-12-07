package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeckCreate(t *testing.T) {
	d := CreateDeck()
	const expected = "2♠,3♠,4♠,5♠,6♠,7♠,8♠,9♠,T♠,J♠,Q♠,K♠,A♠,2♦,3♦,4♦,5♦,6♦,7♦,8♦,9♦,T♦,J♦,Q♦,K♦,A♦,2♣,3♣,4♣,5♣,6♣,7♣,8♣,9♣,T♣,J♣,Q♣,K♣,A♣,2♥,3♥,4♥,5♥,6♥,7♥,8♥,9♥,T♥,J♥,Q♥,K♥,A♥"
	assert.Equal(t, d.String(), expected, "Invalid deck string")
}

func TestDeckDraw(t *testing.T) {
	d := CreateDeck()
	assert.Equal(t, d.Count(), 52, "Invalid draw")
	hand := d.Draw(5)
	assert.Equal(t, len(hand.GetCards()), 5, "Invalid draw")
	assert.Equal(t, d.Count(), 47, "Invalid draw")
	assert.Equal(t, hand.String(), string("2♠,3♠,4♠,5♠,6♠"), "Invalid draw")
}

func TestDeckShuffle(t *testing.T) {
	d := CreateDeck()
	const expected = "2♠,3♠,4♠,5♠,6♠,7♠,8♠,9♠,T♠,J♠,Q♠,K♠,A♠,2♦,3♦,4♦,5♦,6♦,7♦,8♦,9♦,T♦,J♦,Q♦,K♦,A♦,2♣,3♣,4♣,5♣,6♣,7♣,8♣,9♣,T♣,J♣,Q♣,K♣,A♣,2♥,3♥,4♥,5♥,6♥,7♥,8♥,9♥,T♥,J♥,Q♥,K♥,A♥"
	assert.Equal(t, d.String(), expected, "Invalid deck string")
	d.Shuffle(1, DefaultSort)
	assert.NotEqual(t, d.String(), expected, "Invalid deck string")
}
