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

	min, max := card.King, card.Ace
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

	// special case: 10, J, Q, K, A
	if ranks[card.Ten] && ranks[card.Jack] && ranks[card.Queen] && ranks[card.King] && ranks[card.Ace] {
		return true
	}

	return false
}

// rankSortValue returns the sorting value for a rank, treating Ace as highest (14)
func rankSortValue(r card.Rank) int {
	if r == card.Ace {
		return 14
	}
	return int(r)
}

// Sort sorts the hand in ascending order with Ace as the highest card
func (h *Hand) Sort() {
	// Simple bubble sort - sufficient for 5 cards
	for i := 0; i < 4; i++ {
		for j := 0; j < 4-i; j++ {
			if rankSortValue((*h)[j].Rank) > rankSortValue((*h)[j+1].Rank) {
				(*h)[j], (*h)[j+1] = (*h)[j+1], (*h)[j]
			}
		}
	}
}
