package poker

import (
	"sort"

	"github.com/andredupreez/cardgame/game"
)

const (
	HighCard      game.HandRank = 1
	OnePair       game.HandRank = 2
	TwoPair       game.HandRank = 3
	ThreeOfAKind  game.HandRank = 4
	Straight      game.HandRank = 5
	Flush         game.HandRank = 6
	FullHouse     game.HandRank = 7
	FourOfAKind   game.HandRank = 8
	StraightFlush game.HandRank = 9
)

func CreatePoker() *game.Game {
	return game.NewGame(
		5,
		[]game.HandRank{StraightFlush, FourOfAKind, FullHouse, Flush, Straight, ThreeOfAKind, TwoPair, OnePair, HighCard},
		game.MatchTable{
			StraightFlush: game.Match{HandDescription: "Straight Flush", Fun: isStraightFlush},
			FullHouse:     game.Match{HandDescription: "Full House", Fun: isFullHouse},
			Flush:         game.Match{HandDescription: "Flush", Fun: isFlush},
			Straight:      game.Match{HandDescription: "Straight", Fun: isStraight},
			FourOfAKind:   game.Match{HandDescription: "Four of a Kind", Fun: isFourOfAKind},
			ThreeOfAKind:  game.Match{HandDescription: "Three of a Kind", Fun: isThreeOfAKind},
			TwoPair:       game.Match{HandDescription: "Two Pair", Fun: isTwoPair},
			OnePair:       game.Match{HandDescription: "One Pair", Fun: isOnePair},
			HighCard:      game.Match{HandDescription: "High Card", Fun: isHighCard},
		},
	)
}

func isOnePair(groups game.CardGroupMap, hand game.Hand) bool {
	if len(groups) == 4 {
		for _, cards := range groups {
			if len(cards) == 2 {
				return true
			}
		}
	}
	return false
}

func isTwoPair(groups game.CardGroupMap, hand game.Hand) bool {
	if len(groups) == 3 {
		for _, cards := range groups {
			if len(cards) == 3 {
				return false
			}
		}
		return true
	}
	return false
}

func isThreeOfAKind(groups game.CardGroupMap, hand game.Hand) bool {
	if len(groups) == 3 {
		for _, cards := range groups {
			if len(cards) == 3 {
				return true
			}
		}
	}
	return false
}

func isFourOfAKind(groups game.CardGroupMap, hand game.Hand) bool {
	if len(groups) == 2 {
		for _, cards := range groups {
			if len(cards) == 4 {
				return true
			}
		}
	}
	return false
}

func isFullHouse(groups game.CardGroupMap, hand game.Hand) bool {
	if len(groups) == 2 {
		var found2 bool = false
		var found3 bool = false
		for _, cards := range groups {
			if len(cards) == 2 {
				found2 = true
			}
			if len(cards) == 3 {
				found3 = true
			}
		}
		return (found2 && found3)
	}
	return false
}

func isStraight(groups game.CardGroupMap, hand game.Hand) bool {
	cards := hand.GetCards()
	sorted := make([]game.Card, len(cards))
	lastIndex := len(sorted) - 1
	copy(sorted, cards)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].GetFace() < sorted[j].GetFace()
	})
	if len(groups) < len(cards) {
		return false
	}
	if sorted[lastIndex].GetFace() == game.Ace && sorted[0].GetFace() == game.Two && sorted[3].GetFace() == game.Five {
		return true
	}
	if (sorted[lastIndex].GetFace() - sorted[0].GetFace()) == 4 {
		return true
	}
	return false
}

func isFlush(groups game.CardGroupMap, hand game.Hand) bool {
	cards := hand.GetCards()
	suit := cards[0].GetSuite()
	for _, card := range cards {
		if card.GetSuite() != suit {
			return false
		}
	}
	return true
}

func isStraightFlush(groups game.CardGroupMap, hand game.Hand) bool {
	return isStraight(groups, hand) && isFlush(groups, hand)
}

func isHighCard(groups game.CardGroupMap, hand game.Hand) bool {
	return true
}
