package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	TestHand1 HandRank = 1
	TestHand2 HandRank = 2
	TestHand3 HandRank = 3
)

func testhand1(groups CardGroupMap, hand Hand) bool {
	return true
}

func testhand2(groups CardGroupMap, hand Hand) bool {
	return true
}

func testhand3(groups CardGroupMap, hand Hand) bool {
	return false
}

func TestGameDraw(t *testing.T) {
	g := NewGame(4, []HandRank{TestHand1, TestHand2}, MatchTable{
		TestHand1: Match{HandDescription: "Test Hand 1", Fun: testhand1},
		TestHand2: Match{HandDescription: "Test Hand 2", Fun: testhand2},
	})
	hand := g.Draw()
	assert.Equal(t, len(hand.GetCards()), 4, "Incorrect number of cards")
}
func TestGameMatch(t *testing.T) {
	g := NewGame(4,
		[]HandRank{TestHand1, TestHand2, TestHand3}, MatchTable{
			TestHand1: Match{HandDescription: "Test Hand 1", Fun: testhand1},
			TestHand2: Match{HandDescription: "Test Hand 2", Fun: testhand2},
			TestHand3: Match{HandDescription: "Test Hand 3", Fun: testhand3},
		})
	hand := createHand([]string{"AH", "KD", "KS", "TD"})
	rank, desc := g.Match(hand)
	assert.Equal(t, rank, TestHand1, "Incorrect number of cards")
	assert.Equal(t, desc, "Test Hand 1", "Incorrect number of cards")
}

func TestGameNoMatch(t *testing.T) {
	g := NewGame(4,
		[]HandRank{TestHand3, TestHand2, TestHand1}, MatchTable{
			TestHand1: Match{HandDescription: "Test Hand 1", Fun: testhand1},
			TestHand2: Match{HandDescription: "Test Hand 2", Fun: testhand2},
			TestHand3: Match{HandDescription: "Test Hand 3", Fun: testhand3},
		})
	hand := createHand([]string{"AH", "KD", "KS", "TD"})
	rank, desc := g.Match(hand)
	assert.Equal(t, rank, TestHand2, "Incorrect number of cards")
	assert.Equal(t, desc, "Test Hand 2", "Incorrect number of cards")
}

func createHand(handstr []string) Hand {
	results := make([]Card, len(handstr))
	for index, cardstr := range handstr {
		card := StringToCard(cardstr)
		results[index] = card
	}
	return NewHand(results)
}
