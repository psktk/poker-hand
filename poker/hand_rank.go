package poker

import "github.com/psktk/poker-hand/rank"

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
