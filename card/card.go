package card

type Rank uint8
type Suit uint8

const (
	Two Rank = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

const (
	Heart Suit = iota
	Spade
	Diamond
	Club
)

type C struct {
	Rank Rank
	Suit Suit
}

func New(r Rank, s Suit) C {
	return C{r, s}
}
