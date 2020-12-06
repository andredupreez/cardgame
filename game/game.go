package game

type HandRank int

type Match struct {
	HandDescription string
	Fun             func(groups CardGroupMap, hand Hand) bool
}

type MatchTable map[HandRank]Match

type Game struct {
	drawSize   int
	deck       *Deck
	handRank   []HandRank
	matchTable MatchTable
}

func NewGame(drawSize int, handRank []HandRank, matchTable MatchTable) *Game {
	game := &Game{
		deck:       CreateDeck(),
		drawSize:   drawSize,
		handRank:   handRank,
		matchTable: matchTable,
	}
	return game
}

func (g *Game) GetDeck() *Deck {
	return g.deck
}

func (g *Game) Draw() Hand {
	return g.deck.Draw(g.drawSize)
}

func (g *Game) Shuffle() {
	g.deck.Shuffle(1)
}

func (g *Game) Match(hand Hand) (HandRank, string) {
	groups := groupCards(hand)
	for _, rank := range g.handRank {
		match, found := g.matchTable[rank]
		if found && match.Fun(groups, hand) == true {
			return rank, match.HandDescription
		}
	}
	panic("invalid match")
}

func groupCards(hand Hand) CardGroupMap {
	groups := make(CardGroupMap)
	for _, card := range hand.GetCards() {
		groups[card.GetFace()] = append(groups[card.GetFace()], card)
	}
	return groups
}
