package game

import (
	"math/rand"
	"time"
)

type HandRank int

type Match struct {
	HandDescription string
	Fun             func(groups CardGroupMap, hand Hand) bool
}

type MatchTable map[HandRank]Match

type Game struct {
	drawSize   int
	deck       *Deck
	shuffleFun func([]Card) []Card
	handRank   []HandRank
	matchTable MatchTable
}

// initialise a new game with parameters for
//   drawSize: indicates the size of the hands that will be drawed from the deck
//   handRank: This is an array of hand rank order
//   matchTable: A map used to provide match functions for each handRank
//   shuffleFun: A function used to shuffle the deck
func NewGame(drawSize int, handRank []HandRank, matchTable MatchTable, shuffleFun func([]Card) []Card) *Game {
	game := &Game{
		deck:       CreateDeck(),
		drawSize:   drawSize,
		shuffleFun: shuffleFun,
		handRank:   handRank,
		matchTable: matchTable,
	}
	return game
}

// Get the game deck
func (g *Game) GetDeck() *Deck {
	return g.deck
}

// Draw a hand from the deck
func (g *Game) Draw() Hand {
	return g.deck.Draw(g.drawSize)
}

// Shuffle the game deck
func (g *Game) Shuffle() {
	g.deck.Shuffle(1, g.shuffleFun)
}

// Match the hand and rank according to strength
func (g *Game) Match(hand Hand) (HandRank, string) {
	groups := groupCardsByFace(hand)
	for _, rank := range g.handRank {
		match, found := g.matchTable[rank]
		if found && match.Fun(groups, hand) == true {
			return rank, match.HandDescription
		}
	}
	panic("invalid match")
}

// The game default deck sort function
func DefaultSort(cards []Card) []Card {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
	return cards
}

// internal function to group cards by face value
func groupCardsByFace(hand Hand) CardGroupMap {
	groups := make(CardGroupMap)
	for _, card := range hand.GetCards() {
		groups[card.GetFace()] = append(groups[card.GetFace()], card)
	}
	return groups
}
