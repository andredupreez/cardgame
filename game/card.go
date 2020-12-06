package game

import (
	"strings"
)

const MAX_CARDS = 52

type Suit uint8
type Face uint8
type CardGroupMap map[Face][]Card

const (
	Spades   Suit = 1
	Diamonds Suit = 2
	Clubs    Suit = 3
	Hearts   Suit = 4
)

const faces = "23456789TJQKA"
const suits = "SDCH"

var (
	suitUTF8 = []string{"\u2660", "\u2666", "\u2663", "\u2665"}
)

const (
	Two   Face = 1
	Three Face = 2
	Four  Face = 3
	Five  Face = 4
	Six   Face = 5
	Seven Face = 6
	Eight Face = 7
	Nine  Face = 8
	Ten   Face = 9
	Jack  Face = 10
	Queen Face = 11
	King  Face = 12
	Ace   Face = 13
)

type Card struct {
	suit Suit
	face Face
}

func (s Suit) String() string {
	return suits[s-1 : s]
}

func ToSuit(s byte) Suit {
	return Suit(strings.IndexByte(suits, s) + 1)
}

func (s Suit) Symbol() string {
	return suitUTF8[s-1]
}

func (f Face) String() string {
	return faces[f-1 : f]
}

func ToFace(s byte) Face {
	return Face(strings.IndexByte(faces, s) + 1)
}

func (c Card) GetFace() Face {
	return c.face
}

func (c Card) GetSuite() Suit {
	return c.suit
}

func (c Card) String() string {
	return c.face.String() + c.suit.String()
}

func (c Card) ToSymbolString() string {
	return c.face.String() + c.suit.Symbol()
}

func StringToCard(s string) Card {
	if len(s) != 2 {
		panic("Invalid Card String")
	}
	return Card{face: ToFace(s[0]), suit: ToSuit(s[1])}
}
