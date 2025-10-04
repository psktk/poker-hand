package card

import "github.com/psktk/poker-hand/rank"

type Suit uint8

const (
	Heart Suit = iota
	Spade
	Diamond
	Club
)

type C struct {
	Rank rank.Rank
	Suit Suit
}

func New(r rank.Rank, s Suit) C {
	return C{r, s}
}
