package card

import (
	"github.com/psktk/poker-hand/rank"
	"github.com/psktk/poker-hand/suit"
)

var (
	TwoOfHearts   = C{rank.Two, suit.Heart}   // 2♥
	ThreeOfHearts = C{rank.Three, suit.Heart} // 3♥
	FourOfHearts  = C{rank.Four, suit.Heart}  // 4♥
	FiveOfHearts  = C{rank.Five, suit.Heart}  // 5♥
	SixOfHearts   = C{rank.Six, suit.Heart}   // 6♥
	SevenOfHearts = C{rank.Seven, suit.Heart} // 7♥
	EightOfHearts = C{rank.Eight, suit.Heart} // 8♥
	NineOfHearts  = C{rank.Nine, suit.Heart}  // 9♥
	TenOfHearts   = C{rank.Ten, suit.Heart}   // 10♥
	JackOfHearts  = C{rank.Jack, suit.Heart}  // J♥
	QueenOfHearts = C{rank.Queen, suit.Heart} // Q♥
	KingOfHearts  = C{rank.King, suit.Heart}  // K♥
	AceOfHearts   = C{rank.Ace, suit.Heart}   // A♥

	TwoOfSpades   = C{rank.Two, suit.Spade}   // 2♠
	ThreeOfSpades = C{rank.Three, suit.Spade} // 3♠
	FourOfSpades  = C{rank.Four, suit.Spade}  // 4♠
	FiveOfSpades  = C{rank.Five, suit.Spade}  // 5♠
	SixOfSpades   = C{rank.Six, suit.Spade}   // 6♠
	SevenOfSpades = C{rank.Seven, suit.Spade} // 7♠
	EightOfSpades = C{rank.Eight, suit.Spade} // 8♠
	NineOfSpades  = C{rank.Nine, suit.Spade}  // 9♠
	TenOfSpades   = C{rank.Ten, suit.Spade}   // 10♠
	JackOfSpades  = C{rank.Jack, suit.Spade}  // J♠
	QueenOfSpades = C{rank.Queen, suit.Spade} // Q♠
	KingOfSpades  = C{rank.King, suit.Spade}  // K♠
	AceOfSpades   = C{rank.Ace, suit.Spade}   // A♠

	TwoOfDiamonds   = C{rank.Two, suit.Diamond}   // 2♦
	ThreeOfDiamonds = C{rank.Three, suit.Diamond} // 3♦
	FourOfDiamonds  = C{rank.Four, suit.Diamond}  // 4♦
	FiveOfDiamonds  = C{rank.Five, suit.Diamond}  // 5♦
	SixOfDiamonds   = C{rank.Six, suit.Diamond}   // 6♦
	SevenOfDiamonds = C{rank.Seven, suit.Diamond} // 7♦
	EightOfDiamonds = C{rank.Eight, suit.Diamond} // 8♦
	NineOfDiamonds  = C{rank.Nine, suit.Diamond}  // 9♦
	TenOfDiamonds   = C{rank.Ten, suit.Diamond}   // 10♦
	JackOfDiamonds  = C{rank.Jack, suit.Diamond}  // J♦
	QueenOfDiamonds = C{rank.Queen, suit.Diamond} // Q♦
	KingOfDiamonds  = C{rank.King, suit.Diamond}  // K♦
	AceOfDiamonds   = C{rank.Ace, suit.Diamond}   // A♦

	TwoOfClubs   = C{rank.Two, suit.Club}   // 2♣
	ThreeOfClubs = C{rank.Three, suit.Club} // 3♣
	FourOfClubs  = C{rank.Four, suit.Club}  // 4♣
	FiveOfClubs  = C{rank.Five, suit.Club}  // 5♣
	SixOfClubs   = C{rank.Six, suit.Club}   // 6♣
	SevenOfClubs = C{rank.Seven, suit.Club} // 7♣
	EightOfClubs = C{rank.Eight, suit.Club} // 8♣
	NineOfClubs  = C{rank.Nine, suit.Club}  // 9♣
	TenOfClubs   = C{rank.Ten, suit.Club}   // 10♣
	JackOfClubs  = C{rank.Jack, suit.Club}  // J♣
	QueenOfClubs = C{rank.Queen, suit.Club} // Q♣
	KingOfClubs  = C{rank.King, suit.Club}  // K♣
	AceOfClubs   = C{rank.Ace, suit.Club}   // A♣
)

type C struct {
	Rank rank.R
	Suit suit.S
}

func New(r rank.R, s suit.S) C {
	return C{r, s}
}

// Compare returns 1 if c > other, -1 if c < other, 0 if equal
func (c C) Compare(other C) int {
	if c.Rank > other.Rank {
		return 1
	}
	if c.Rank < other.Rank {
		return -1
	}
	return 0
}
