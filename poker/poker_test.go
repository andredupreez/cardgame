package poker

import (
	"testing"

	"github.com/andredupreez/cardgame/game"
	"github.com/stretchr/testify/assert"
)

func TestOnePair(t *testing.T) {
	g := CreatePoker()
	h1 := createHand([]string{"AH", "AS", "2S", "5S", "9C"})
	hand, description := g.Match(h1)
	assert.Equal(t, hand, OnePair, "The hand should be a One Pair")
	assert.Equal(t, description, "One Pair", "The hand should be a One Pair")
}

func TestTwoPair(t *testing.T) {
	g := CreatePoker()
	h1 := createHand([]string{"AH", "AS", "2S", "2C", "9C"})
	hand, description := g.Match(h1)
	assert.Equal(t, hand, TwoPair, "The hand should be a Two Pair")
	assert.Equal(t, description, "Two Pair", "The hand should be a Two Pair")
}

func TestThreeOfAKind(t *testing.T) {
	g := CreatePoker()
	h1 := createHand([]string{"AH", "AS", "AC", "2C", "9C"})
	hand, description := g.Match(h1)
	assert.Equal(t, hand, ThreeOfAKind, "The hand should be a Three of a Kind")
	assert.Equal(t, description, "Three of a Kind", "The hand should be a Three of a Kind")
}

func TestFourOfAKind(t *testing.T) {
	g := CreatePoker()
	h1 := createHand([]string{"AH", "AS", "AC", "AD", "9C"})
	hand, description := g.Match(h1)
	assert.Equal(t, hand, FourOfAKind, "The hand should be a Four of a Kind")
	assert.Equal(t, description, "Four of a Kind", "The hand should be a Four of a Kind")
}

func TestStraightAceHigh(t *testing.T) {
	g := CreatePoker()
	h1 := createHand([]string{"AH", "KH", "QS", "JS", "TC"})
	hand, description := g.Match(h1)
	assert.Equal(t, hand, Straight, "The hand should be a straight")
	assert.Equal(t, description, "Straight", "The hand should be a straight")
}

func TestStraightAceLow(t *testing.T) {
	g := CreatePoker()
	h1 := createHand([]string{"AH", "2S", "3H", "4C", "5C"})
	hand, description := g.Match(h1)
	assert.Equal(t, hand, Straight, "The hand should be a straight")
	assert.Equal(t, description, "Straight", "The hand should be a straight")
}

func TestStraightMiddleCards(t *testing.T) {
	g := CreatePoker()
	h1 := createHand([]string{"3S", "4H", "5C", "6C", "7H"})
	hand, description := g.Match(h1)
	assert.Equal(t, hand, Straight, "The hand should be a straight")
	assert.Equal(t, description, "Straight", "The hand should be a straight")
}

func TestFlush(t *testing.T) {
	g := CreatePoker()
	h1 := createHand([]string{"3S", "4S", "5S", "6S", "TS"})
	hand, description := g.Match(h1)
	assert.Equal(t, hand, Flush, "The hand should be a flush")
	assert.Equal(t, description, "Flush", "The hand should be a flush")
}

func TestFullHouse(t *testing.T) {
	g := CreatePoker()
	h1 := createHand([]string{"3S", "3C", "3H", "6S", "6H"})
	hand, description := g.Match(h1)
	assert.Equal(t, hand, FullHouse, "The hand should be a Full House")
	assert.Equal(t, description, "Full House", "The hand should be a Full House")
}

func TestStraightFlush(t *testing.T) {
	g := CreatePoker()
	h1 := createHand([]string{"2S", "3S", "4S", "5S", "6S"})
	hand, description := g.Match(h1)
	assert.Equal(t, hand, StraightFlush, "The hand should be a Straight Flush")
	assert.Equal(t, description, "Straight Flush", "The hand should be a Straight Flush")
}

//---------------------------------------------------------------------------------
// Internal functions
//---------------------------------------------------------------------------------

func createHand(handstr []string) game.Hand {
	results := make([]game.Card, len(handstr))
	for index, cardstr := range handstr {
		card := game.StringToCard(cardstr)
		results[index] = card
	}
	return game.NewHand(results)
}
