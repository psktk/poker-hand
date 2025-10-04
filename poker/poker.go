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

	if len(ranks) < 5 {
		return false
	}

	// copy hand to sort without modifying original
	h2 := h
	h2.Sort()
	if h2[4].Rank-h2[0].Rank == 4 {
		return true
	}

	// special case: A, 2, 3, 4, 5
	if ranks[card.Ace] && ranks[card.Two] && ranks[card.Three] && ranks[card.Four] && ranks[card.Five] {
		return true
	}

	return false
}

func (h Hand) IsStraightFlush() bool {
	return h.IsFlush() && h.IsStraight()
}

func (h *Hand) Sort() {
	for i := range 4 {
		for j := range 4 - i {
			if (*h)[j].Rank > (*h)[j+1].Rank {
				(*h)[j], (*h)[j+1] = (*h)[j+1], (*h)[j]
			}
		}
	}
}
