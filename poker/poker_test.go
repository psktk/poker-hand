package poker

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/psktk/poker-hand/card"
	"github.com/psktk/poker-hand/rank"
)

func TestCompareStraight(t *testing.T) {
	t.Run("should return 1 if this hand's straight is higher", func(t *testing.T) {
		h1 := Hand{card.NineOfHearts, card.TenOfSpades, card.JackOfDiamonds, card.QueenOfClubs, card.KingOfHearts}
		h2 := Hand{card.EightOfHearts, card.NineOfDiamonds, card.TenOfClubs, card.JackOfSpades, card.QueenOfHearts}
		assert.Equal(t, 1, h1.compareStraight(h2))
	})

	t.Run("should return -1 if this hand's straight is lower", func(t *testing.T) {
		h1 := Hand{card.EightOfHearts, card.NineOfDiamonds, card.TenOfClubs, card.JackOfSpades, card.QueenOfHearts}
		h2 := Hand{card.NineOfHearts, card.TenOfSpades, card.JackOfDiamonds, card.QueenOfClubs, card.KingOfHearts}
		assert.Equal(t, -1, h1.compareStraight(h2))
	})

	t.Run("should return 0 if this hand's straight is the same", func(t *testing.T) {
		h1 := Hand{card.NineOfHearts, card.TenOfSpades, card.JackOfDiamonds, card.QueenOfClubs, card.KingOfHearts}
		h2 := Hand{card.NineOfDiamonds, card.TenOfClubs, card.JackOfSpades, card.QueenOfHearts, card.KingOfHearts}
		assert.Equal(t, 0, h1.compareStraight(h2))
	})

	t.Run("should return -1 if the first hand is a low straight and the second is not", func(t *testing.T) {
		h1 := Hand{card.AceOfHearts, card.TwoOfSpades, card.ThreeOfDiamonds, card.FourOfClubs, card.FiveOfHearts}
		h2 := Hand{card.TwoOfDiamonds, card.ThreeOfClubs, card.FourOfSpades, card.FiveOfHearts, card.SixOfHearts}
		assert.Equal(t, -1, h1.compareStraight(h2))
	})

	t.Run("should return 1 if the first hand is not a low straight and the second is", func(t *testing.T) {
		h1 := Hand{card.TwoOfDiamonds, card.ThreeOfClubs, card.FourOfSpades, card.FiveOfHearts, card.SixOfHearts}
		h2 := Hand{card.AceOfHearts, card.TwoOfSpades, card.ThreeOfDiamonds, card.FourOfClubs, card.FiveOfHearts}
		assert.Equal(t, 1, h1.compareStraight(h2))
	})

	t.Run("should return 0 if both hands are low straights", func(t *testing.T) {
		h1 := Hand{card.AceOfHearts, card.TwoOfSpades, card.ThreeOfDiamonds, card.FourOfClubs, card.FiveOfHearts}
		h2 := Hand{card.AceOfDiamonds, card.TwoOfClubs, card.ThreeOfSpades, card.FourOfHearts, card.FiveOfHearts}
		assert.Equal(t, 0, h1.compareStraight(h2))
	})
}

func TestCompareFlush(t *testing.T) {
	t.Run("should return 1 if this hand's flush is higher", func(t *testing.T) {
		h1 := Hand{card.TwoOfSpades, card.FourOfSpades, card.SixOfSpades, card.EightOfSpades, card.AceOfSpades}
		h2 := Hand{card.TwoOfHearts, card.FourOfHearts, card.SixOfHearts, card.EightOfHearts, card.KingOfHearts}
		assert.Equal(t, 1, h1.compareFlush(h2))
	})

	t.Run("should return -1 if this hand's flush is lower", func(t *testing.T) {
		h1 := Hand{card.TwoOfHearts, card.FourOfHearts, card.SixOfHearts, card.EightOfHearts, card.KingOfHearts}
		h2 := Hand{card.TwoOfSpades, card.FourOfSpades, card.SixOfSpades, card.EightOfSpades, card.AceOfSpades}
		assert.Equal(t, -1, h1.compareFlush(h2))
	})

	t.Run("should return 0 if this hand's flush is the same", func(t *testing.T) {
		h1 := Hand{card.TwoOfHearts, card.FourOfHearts, card.SixOfHearts, card.EightOfHearts, card.TenOfHearts}
		h2 := Hand{card.TwoOfDiamonds, card.FourOfDiamonds, card.SixOfDiamonds, card.EightOfDiamonds, card.TenOfDiamonds}
		assert.Equal(t, 0, h1.compareFlush(h2))
	})
}

func TestQuadRank(t *testing.T) {
	t.Run("should return the rank of the four of a kind", func(t *testing.T) {
		q, _ := Hand{card.AceOfHearts, card.AceOfSpades, card.AceOfDiamonds, card.AceOfClubs, card.FiveOfHearts}.quadRank()
		assert.Equal(t, rank.Ace, q)
	})

	t.Run("should return the rank of the kicker", func(t *testing.T) {
		_, k := Hand{card.AceOfHearts, card.AceOfSpades, card.AceOfDiamonds, card.AceOfClubs, card.FiveOfHearts}.quadRank()
		assert.Equal(t, rank.Five, k)
	})
}

func TestCompareFourOfAKind(t *testing.T) {
	t.Run("should return 1 if this hand's four of a kind is higher", func(t *testing.T) {
		h1 := Hand{card.AceOfHearts, card.AceOfSpades, card.AceOfDiamonds, card.AceOfClubs, card.FiveOfHearts}
		h2 := Hand{card.KingOfHearts, card.KingOfSpades, card.KingOfDiamonds, card.KingOfClubs, card.FiveOfHearts}
		assert.Equal(t, 1, h1.compareFourOfAKind(h2))
	})

	t.Run("should return -1 if this hand's four of a kind is lower", func(t *testing.T) {
		h1 := Hand{card.KingOfHearts, card.KingOfSpades, card.KingOfDiamonds, card.KingOfClubs, card.FiveOfHearts}
		h2 := Hand{card.AceOfHearts, card.AceOfSpades, card.AceOfDiamonds, card.AceOfClubs, card.FiveOfHearts}
		assert.Equal(t, -1, h1.compareFourOfAKind(h2))
	})

	t.Run("should return 1 if this hand's kicker is higher", func(t *testing.T) {
		h1 := Hand{card.AceOfHearts, card.AceOfSpades, card.AceOfDiamonds, card.AceOfClubs, card.SixOfHearts}
		h2 := Hand{card.AceOfHearts, card.AceOfSpades, card.AceOfDiamonds, card.AceOfClubs, card.FiveOfHearts}
		assert.Equal(t, 1, h1.compareFourOfAKind(h2))
	})

	t.Run("should return -1 if this hand's kicker is lower", func(t *testing.T) {
		h1 := Hand{card.AceOfHearts, card.AceOfSpades, card.AceOfDiamonds, card.AceOfClubs, card.FiveOfHearts}
		h2 := Hand{card.AceOfHearts, card.AceOfSpades, card.AceOfDiamonds, card.AceOfClubs, card.SixOfHearts}
		assert.Equal(t, -1, h1.compareFourOfAKind(h2))
	})

	t.Run("should return 0 if both hands are the same", func(t *testing.T) {
		h1 := Hand{card.AceOfHearts, card.AceOfSpades, card.AceOfDiamonds, card.AceOfClubs, card.FiveOfHearts}
		assert.Equal(t, 0, h1.compareFourOfAKind(h1))
	})
}

func TestIsLowStraight(t *testing.T) {
	t.Run("should return true for a ace-low straight hand", func(t *testing.T) {
		h := Hand{card.AceOfHearts, card.TwoOfSpades, card.ThreeOfDiamonds, card.FourOfClubs, card.FiveOfHearts}
		assert.True(t, h.isLowStraight())
	})

	t.Run("should return false for a ace-high straight hand", func(t *testing.T) {
		h := Hand{card.TenOfHearts, card.JackOfSpades, card.QueenOfDiamonds, card.KingOfClubs, card.AceOfHearts}
		assert.False(t, h.isLowStraight())
	})
}

func TestSort(t *testing.T) {
	t.Run("should sort the hand in ascending order by rank", func(t *testing.T) {
		h := NewHand(
			card.KingOfHearts,
			card.ThreeOfSpades,
			card.QueenOfDiamonds,
			card.FiveOfClubs,
			card.TenOfHearts,
		)
		h.Sort()
		expected := NewHand(
			card.ThreeOfSpades,
			card.FiveOfClubs,
			card.TenOfHearts,
			card.QueenOfDiamonds,
			card.KingOfHearts,
		)
		assert.Equal(t, expected, h)
	})

	t.Run("should sort a hand with duplicate ranks", func(t *testing.T) {
		h := NewHand(
			card.FiveOfHearts,
			card.ThreeOfSpades,
			card.FiveOfDiamonds,
			card.TwoOfClubs,
			card.ThreeOfHearts,
		)
		h.Sort()
		expected := NewHand(
			card.TwoOfClubs,
			card.ThreeOfSpades,
			card.ThreeOfHearts,
			card.FiveOfHearts,
			card.FiveOfDiamonds,
		)
		assert.Equal(t, expected, h)
	})

	t.Run("should sort a hand with ace as the highest rank", func(t *testing.T) {
		h := NewHand(
			card.ThreeOfHearts,
			card.AceOfSpades,
			card.TenOfDiamonds,
			card.FourOfClubs,
			card.KingOfHearts,
		)
		h.Sort()
		expected := NewHand(
			card.ThreeOfHearts,
			card.FourOfClubs,
			card.TenOfDiamonds,
			card.KingOfHearts,
			card.AceOfSpades,
		)
		assert.Equal(t, expected, h)
	})
}

func TestCountPairs(t *testing.T) {
	t.Run("should return 0 for a hand with no pairs", func(t *testing.T) {
		h := NewHand(
			card.AceOfHearts,
			card.KingOfSpades,
			card.QueenOfDiamonds,
			card.JackOfClubs,
			card.TenOfHearts,
		)
		assert.Equal(t, 0, h.countPairs())
	})

	t.Run("should return 1 for a hand with one pair", func(t *testing.T) {
		h := NewHand(
			card.AceOfHearts,
			card.AceOfSpades,
			card.QueenOfDiamonds,
			card.JackOfClubs,
			card.TenOfHearts,
		)
		assert.Equal(t, 1, h.countPairs())
	})

	t.Run("should return 2 for a hand with two pairs", func(t *testing.T) {
		h := NewHand(
			card.AceOfHearts,
			card.AceOfSpades,
			card.KingOfDiamonds,
			card.KingOfClubs,
			card.TenOfHearts,
		)
		assert.Equal(t, 2, h.countPairs())
	})

	t.Run("should return 1 for a full house hand", func(t *testing.T) {
		h := NewHand(
			card.AceOfHearts,
			card.AceOfSpades,
			card.KingOfDiamonds,
			card.KingOfClubs,
			card.KingOfHearts,
		)
		assert.Equal(t, 1, h.countPairs())
	})
}
