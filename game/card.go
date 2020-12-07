package game

import (
	"strings"
)

const MAX_CARDS = 52

type Suit uint8
type Face uint8
type CardGroupMap map[Face][]Card

const (
	Spades   Suit = 1 // "S"
	Diamonds Suit = 2 // "D"
	Clubs    Suit = 3 // "C"
	Hearts   Suit = 4 // "H"
)

const faces = "23456789TJQKA"
const suits = "SDCH"

var (
	suitUTF8 = []string{"\u2660", "\u2666", "\u2663", "\u2665"}
)

const (
	Two   Face = 1  // "2"
	Three Face = 2  // "3"
	Four  Face = 3  // "4"
	Five  Face = 4  // "5"
	Six   Face = 5  // "6"
	Seven Face = 6  // "7"
	Eight Face = 7  // "8"
	Nine  Face = 8  // "9"
	Ten   Face = 9  // "T"
	Jack  Face = 10 // "J"
	Queen Face = 11 // "Q"
	King  Face = 12 // "K"
	Ace   Face = 13 // "A"
)

type Card struct {
	suit Suit
	face Face
}

// Get the Suit string description
func (s Suit) String() string {
	return suits[s-1 : s]
}

// Get the Suit symbol description
func (s Suit) Symbol() string {
	return suitUTF8[s-1]
}

// Get the Face string description
func (f Face) String() string {
	return faces[f-1 : f]
}

// Get the Suit value for the specified string character
func ToSuit(s byte) Suit {
	return Suit(strings.IndexByte(suits, s) + 1)
}

// Get the Face value for the specified string character
func ToFace(s byte) Face {
	return Face(strings.IndexByte(faces, s) + 1)
}

// Get the Card Face value
func (c Card) GetFace() Face {
	return c.face
}

// Get the Card Suit value
func (c Card) GetSuit() Suit {
	return c.suit
}

// Get the card description
func (c Card) String() string {
	return c.face.String() + c.suit.String()
}

// Get the card description using suit symbols
func (c Card) ToSymbolString() string {
	return c.face.String() + c.suit.Symbol()
}

// Convert a String "AH" card format to a Card struct
func StringToCard(s string) Card {
	if len(s) != 2 {
		panic("invalid card string")
	}
	return Card{face: ToFace(s[0]), suit: ToSuit(s[1])}
}
