package card

import (
	"github.com/psktk/poker-hand/rank"
	"github.com/psktk/poker-hand/suit"
)

type C struct {
	Rank rank.R
	Suit suit.S
}

func New(r rank.R, s suit.S) C {
	return C{r, s}
}
