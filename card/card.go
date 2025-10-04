package card

import (
	"github.com/psktk/poker-hand/rank"
	"github.com/psktk/poker-hand/suit"
)

type C struct {
	Rank rank.Rank
	Suit suit.Suit
}

func New(r rank.Rank, s suit.Suit) C {
	return C{r, s}
}
