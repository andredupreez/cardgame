package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCardDefault(t *testing.T) {
	c := Card{}
	assert.Equal(t, c.face, Face(0), "Invalid Face")
	assert.Equal(t, c.suit, Suit(0), "Invalid Suit")
}
func TestCardSuit(t *testing.T) {
	assert.Equal(t, Card{suit: Hearts}.GetSuite().String(), "H", "Invalid Suit")
	assert.Equal(t, Card{suit: Diamonds}.GetSuite().String(), "D", "Invalid Suit")
	assert.Equal(t, Card{suit: Spades}.GetSuite().String(), "S", "Invalid Suit")
	assert.Equal(t, Card{suit: Clubs}.GetSuite().String(), "C", "Invalid Suit")
}
func TestCardSuitSymbol(t *testing.T) {
	assert.Equal(t, Card{suit: Hearts}.GetSuite().Symbol(), "♥", "Invalid Suit Symbol")
	assert.Equal(t, Card{suit: Diamonds}.GetSuite().Symbol(), "♦", "Invalid Suit Symbol")
	assert.Equal(t, Card{suit: Spades}.GetSuite().Symbol(), "♠", "Invalid Suit Symbol")
	assert.Equal(t, Card{suit: Clubs}.GetSuite().Symbol(), "♣", "Invalid Suit Symbol")
}
func TestCardFace(t *testing.T) {
	assert.Equal(t, Card{face: Two}.GetFace().String(), "2", "Invalid Face")
	assert.Equal(t, Card{face: Three}.GetFace().String(), "3", "Invalid Face")
	assert.Equal(t, Card{face: Four}.GetFace().String(), "4", "Invalid Face")
	assert.Equal(t, Card{face: Five}.GetFace().String(), "5", "Invalid Face")
	assert.Equal(t, Card{face: Six}.GetFace().String(), "6", "Invalid Face")
	assert.Equal(t, Card{face: Seven}.GetFace().String(), "7", "Invalid Face")
	assert.Equal(t, Card{face: Eight}.GetFace().String(), "8", "Invalid Face")
	assert.Equal(t, Card{face: Nine}.GetFace().String(), "9", "Invalid Face")
	assert.Equal(t, Card{face: Ten}.GetFace().String(), "T", "Invalid Face")
	assert.Equal(t, Card{face: Jack}.GetFace().String(), "J", "Invalid Face")
	assert.Equal(t, Card{face: Queen}.GetFace().String(), "Q", "Invalid Face")
	assert.Equal(t, Card{face: King}.GetFace().String(), "K", "Invalid Face")
	assert.Equal(t, Card{face: Ace}.GetFace().String(), "A", "Invalid Face")
}
func TestCardString(t *testing.T) {
	assert.Equal(t, Card{suit: Hearts, face: Two}.String(), "2H", "Invalid Card string")
	assert.Equal(t, Card{suit: Diamonds, face: King}.String(), "KD", "Invalid Card string")
	assert.Equal(t, Card{suit: Spades, face: Ace}.String(), "AS", "Invalid Card string")
	assert.Equal(t, Card{suit: Clubs, face: Ten}.String(), "TC", "Invalid Card string")
}
func TestStringToCard(t *testing.T) {
	assert.Equal(t, StringToCard("2H"), Card{suit: Hearts, face: Two}, "Invalid Card string")
	assert.Equal(t, StringToCard("KD"), Card{suit: Diamonds, face: King}, "Invalid Card string")
	assert.Equal(t, StringToCard("AS"), Card{suit: Spades, face: Ace}, "Invalid Card string")
	assert.Equal(t, StringToCard("TC"), Card{suit: Clubs, face: Ten}, "Invalid Card string")
}
