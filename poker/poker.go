package poker

import (
	"github.com/psktk/poker-hand/card"
	"github.com/psktk/poker-hand/rank"
)

type Hand [5]card.C

func NewHand(c1, c2, c3, c4, c5 card.C) Hand {
	return Hand{c1, c2, c3, c4, c5}
}

func (h Hand) compareStraight(other Hand) int {
	if h.isLowStraight() {
		h = Hand{h[4], h[0], h[1], h[2], h[3]}
	}
	if other.isLowStraight() {
		other = Hand{other[4], other[0], other[1], other[2], other[3]}
	}
	return h[4].Compare(other[4])
}

func (h Hand) compareFlush(other Hand) int {
	h.Sort()
	other.Sort()
	for i := 4; i >= 0; i-- {
		if cmp := h[i].Compare(other[i]); cmp != 0 {
			return cmp
		}
	}
	return 0
}

func (h Hand) quadRank() (q, k rank.R) {
	for r, count := range h.rankCountMap() {
		switch count {
		case 4:
			q = r
		case 1:
			k = r
		}
	}
	return
}

func (h Hand) compareFourOfAKind(other Hand) int {
	q1, k1 := h.quadRank()
	q2, k2 := other.quadRank()
	if cmp := q1.Compare(q2); cmp != 0 {
		return cmp
	}
	return k1.Compare(k2)
}

func (h Hand) isLowStraight() bool {
	h.Sort()
	return h[0].Rank == rank.Two &&
		h[1].Rank == rank.Three &&
		h[2].Rank == rank.Four &&
		h[3].Rank == rank.Five &&
		h[4].Rank == rank.Ace
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
