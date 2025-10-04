package poker

import (
	"github.com/psktk/poker-hand/card"
)

type Hand [5]card.C

func NewHand(c1, c2, c3, c4, c5 card.C) Hand {
	return Hand{c1, c2, c3, c4, c5}
}

func (h Hand) IsFlush() bool {
	suit := h[0].Suit
	for _, c := range h {
		if c.Suit != suit {
			return false
		}
	}
	return true
}

func (h Hand) IsStraight() bool {
	ranks := make(map[card.Rank]bool)
	for _, c := range h {
		ranks[c.Rank] = true
	}

	min, max := card.Ace, card.Two
	for r := range ranks {
		if r < min {
			min = r
		}
		if r > max {
			max = r
		}
	}

	if len(ranks) < 5 {
		return false
	}

	if max-min == 4 {
		return true
	}

	// special case: A, 2, 3, 4, 5
	if ranks[card.Ace] && ranks[card.Two] && ranks[card.Three] && ranks[card.Four] && ranks[card.Five] {
		return true
	}

	return false
}

func (h *Hand) Sort() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4-i; j++ {
			if (*h)[j].Rank > (*h)[j+1].Rank {
				(*h)[j], (*h)[j+1] = (*h)[j+1], (*h)[j]
			}
		}
	}
}
