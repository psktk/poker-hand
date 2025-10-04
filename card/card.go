package card

type Rank uint8
type Suit uint8

const (
	Ace Rank = iota
	Two
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
)

const (
	Heart Suit = iota
	Spade
	Diamond
	Club
)

type Card struct {
	Rank Rank
	Suit Suit
}

func New(r Rank, s Suit) Card {
	return Card{r, s}
}
