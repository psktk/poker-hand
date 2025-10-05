package poker

import (
	"github.com/psktk/poker-hand/card"
	"github.com/psktk/poker-hand/rank"
)

type HandRank uint8

const (
	HighCard HandRank = iota
	OnePair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
	RoyalFlush
)

type Hand [5]card.C

func NewHand(c1, c2, c3, c4, c5 card.C) Hand {
	return Hand{c1, c2, c3, c4, c5}
}

func (h Hand) Rank() HandRank {
	switch {
	case h.IsRoyalFlush():
		return RoyalFlush
	case h.IsStraightFlush():
		return StraightFlush
	case h.IsFourOfAKind():
		return FourOfAKind
	case h.IsFullHouse():
		return FullHouse
	case h.IsFlush():
		return Flush
	case h.IsStraight():
		return Straight
	case h.IsThreeOfAKind():
		return ThreeOfAKind
	case h.IsTwoPair():
		return TwoPair
	case h.IsOnePair():
		return OnePair
	default:
		return HighCard
	}
}

func (h Hand) compareStraight(other Hand) int {
	h.Sort()
	other.Sort()
	return h[4].Compare(other[4])
}

func (h Hand) isLowStraight() bool {
	h.Sort()
	return h[0].Rank == rank.Two &&
		h[1].Rank == rank.Three &&
		h[2].Rank == rank.Four &&
		h[3].Rank == rank.Five &&
		h[4].Rank == rank.Ace
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
	ranks := make(map[rank.R]bool)
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
	if ranks[rank.Ace] && ranks[rank.Two] && ranks[rank.Three] && ranks[rank.Four] && ranks[rank.Five] {
		return true
	}

	return false
}

func (h Hand) IsStraightFlush() bool {
	return h.IsFlush() && h.IsStraight()
}

func (h Hand) IsRoyalFlush() bool {
	h2 := h
	h2.Sort()
	return h.IsStraightFlush() && h2[4].Rank == rank.Ace
}

func (h Hand) IsFourOfAKind() bool {
	for _, count := range h.rankCountMap() {
		if count == 4 {
			return true
		}
	}
	return false
}

func (h Hand) IsFullHouse() bool {
	return h.hasTriplet() && h.countPairs() == 1
}

func (h Hand) IsThreeOfAKind() bool {
	return h.hasTriplet() && h.countPairs() == 0
}

func (h Hand) IsTwoPair() bool {
	return h.countPairs() == 2
}

func (h Hand) IsOnePair() bool {
	return h.countPairs() == 1 && !h.hasTriplet()
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

func (h Hand) rankCountMap() map[rank.R]int {
	m := make(map[rank.R]int)
	for _, c := range h {
		m[c.Rank]++
	}
	return m
}

func (h Hand) hasTriplet() bool {
	for _, count := range h.rankCountMap() {
		if count == 3 {
			return true
		}
	}
	return false
}

func (h Hand) countPairs() (pairs int) {
	for _, count := range h.rankCountMap() {
		if count == 2 {
			pairs++
		}
	}
	return
}
