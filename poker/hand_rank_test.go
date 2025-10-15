package poker

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/psktk/poker-hand/card"
)

func TestRank(t *testing.T) {
	t.Run("should return RoyalFlush for a royal flush hand", func(t *testing.T) {
		h := Hand{card.TenOfHearts, card.JackOfHearts, card.QueenOfHearts, card.KingOfHearts, card.AceOfHearts}
		assert.Equal(t, RoyalFlush, h.Rank())
	})

	t.Run("should return StraightFlush for a straight flush hand", func(t *testing.T) {
		h := Hand{card.NineOfHearts, card.TenOfHearts, card.JackOfHearts, card.QueenOfHearts, card.KingOfHearts}
		assert.Equal(t, StraightFlush, h.Rank())
	})

	t.Run("should return FourOfAKind for a four of a kind hand", func(t *testing.T) {
		h := Hand{card.AceOfHearts, card.AceOfSpades, card.AceOfDiamonds, card.AceOfClubs, card.FiveOfHearts}
		assert.Equal(t, FourOfAKind, h.Rank())
	})

	t.Run("should return FullHouse for a full house hand", func(t *testing.T) {
		h := Hand{card.AceOfHearts, card.AceOfSpades, card.AceOfDiamonds, card.KingOfClubs, card.KingOfHearts}
		assert.Equal(t, FullHouse, h.Rank())
	})

	t.Run("should return Flush for a flush hand", func(t *testing.T) {
		h := Hand{card.TwoOfSpades, card.FourOfSpades, card.SixOfSpades, card.EightOfSpades, card.TenOfSpades}
		assert.Equal(t, Flush, h.Rank())
	})

	t.Run("should return Straight for a straight hand", func(t *testing.T) {
		h := Hand{card.NineOfHearts, card.TenOfSpades, card.JackOfDiamonds, card.QueenOfClubs, card.KingOfHearts}
		assert.Equal(t, Straight, h.Rank())
	})

	t.Run("should return ThreeOfAKind for a three of a kind hand", func(t *testing.T) {
		h := Hand{card.AceOfHearts, card.AceOfSpades, card.AceOfDiamonds, card.KingOfClubs, card.FiveOfHearts}
		assert.Equal(t, ThreeOfAKind, h.Rank())
	})

	t.Run("should return TwoPair for a two pair hand", func(t *testing.T) {
		h := Hand{card.AceOfHearts, card.AceOfSpades, card.KingOfDiamonds, card.KingOfClubs, card.FiveOfHearts}
		assert.Equal(t, TwoPair, h.Rank())
	})

	t.Run("should return OnePair for a one pair hand", func(t *testing.T) {
		h := Hand{card.AceOfHearts, card.AceOfSpades, card.QueenOfDiamonds, card.JackOfClubs, card.TenOfHearts}
		assert.Equal(t, OnePair, h.Rank())
	})

	t.Run("should return HighCard for a high card hand", func(t *testing.T) {
		h := Hand{card.AceOfHearts, card.KingOfSpades, card.QueenOfDiamonds, card.JackOfClubs, card.NineOfClubs}
		assert.Equal(t, HighCard, h.Rank())
	})
}

func TestIsFlush(t *testing.T) {
	t.Run("should return true for a straight flush hand", func(t *testing.T) {
		h := NewHand(
			card.AceOfHearts,
			card.TwoOfHearts,
			card.ThreeOfHearts,
			card.FourOfHearts,
			card.FiveOfHearts,
		)
		assert.True(t, h.IsFlush())
	})

	t.Run("should return true for a flush hand", func(t *testing.T) {
		h := NewHand(
			card.TwoOfSpades,
			card.FourOfSpades,
			card.SixOfSpades,
			card.EightOfSpades,
			card.TenOfSpades,
		)
		assert.True(t, h.IsFlush())
	})

	t.Run("should return false for a non-flush hand", func(t *testing.T) {
		h := NewHand(
			card.TwoOfHearts,
			card.FourOfHearts,
			card.SixOfHearts,
			card.EightOfHearts,
			card.TenOfSpades,
		)
		assert.False(t, h.IsFlush())
	})
}

func TestIsStraight(t *testing.T) {
	t.Run("should return true for a straight hand two to six", func(t *testing.T) {
		h := NewHand(
			card.AceOfHearts,
			card.TwoOfSpades,
			card.ThreeOfDiamonds,
			card.FourOfClubs,
			card.FiveOfHearts,
		)
		assert.True(t, h.IsStraight())
	})

	t.Run("should return true for a straight hand nine to king", func(t *testing.T) {
		h := NewHand(
			card.NineOfHearts,
			card.TenOfSpades,
			card.JackOfDiamonds,
			card.QueenOfClubs,
			card.KingOfHearts,
		)
		assert.True(t, h.IsStraight())
	})

	t.Run("should return false for a hand with duplicate ranks", func(t *testing.T) {
		h := NewHand(
			card.TwoOfHearts,
			card.ThreeOfSpades,
			card.ThreeOfDiamonds,
			card.FourOfClubs,
			card.FiveOfHearts,
		)
		assert.False(t, h.IsStraight())
	})

	t.Run("should return false for a non-straight hand", func(t *testing.T) {
		h := NewHand(
			card.TwoOfHearts,
			card.FourOfSpades,
			card.SixOfDiamonds,
			card.EightOfClubs,
			card.TenOfHearts,
		)
		assert.False(t, h.IsStraight())
	})

	t.Run("should return true for a straight hand ten to ace", func(t *testing.T) {
		h := NewHand(
			card.TenOfHearts,
			card.JackOfSpades,
			card.QueenOfDiamonds,
			card.KingOfClubs,
			card.AceOfHearts,
		)
		assert.True(t, h.IsStraight())
	})

	t.Run("should return true for a straight hand ace to five", func(t *testing.T) {
		h := NewHand(
			card.AceOfHearts,
			card.TwoOfSpades,
			card.ThreeOfDiamonds,
			card.FourOfClubs,
			card.FiveOfHearts,
		)
		assert.True(t, h.IsStraight())
	})
}

func TestIsStraightFlush(t *testing.T) {
	t.Run("should return true for a straight flush hand", func(t *testing.T) {
		h := NewHand(
			card.AceOfHearts,
			card.TwoOfHearts,
			card.ThreeOfHearts,
			card.FourOfHearts,
			card.FiveOfHearts,
		)
		assert.True(t, h.IsStraightFlush())
	})

	t.Run("should return false for a flush hand", func(t *testing.T) {
		h := NewHand(
			card.TwoOfSpades,
			card.FourOfSpades,
			card.SixOfSpades,
			card.EightOfSpades,
			card.TenOfSpades,
		)
		assert.False(t, h.IsStraightFlush())
	})

	t.Run("should return false for a straight hand", func(t *testing.T) {
		h := NewHand(
			card.TwoOfHearts,
			card.ThreeOfSpades,
			card.FourOfDiamonds,
			card.FiveOfClubs,
			card.SixOfHearts,
		)
		assert.False(t, h.IsStraightFlush())
	})

	t.Run("should return false for a non-straight flush hand", func(t *testing.T) {
		h := NewHand(
			card.TwoOfHearts,
			card.FourOfHearts,
			card.SixOfHearts,
			card.EightOfHearts,
			card.TenOfSpades,
		)
		assert.False(t, h.IsStraightFlush())
	})
}

func TestIsRoyalFlush(t *testing.T) {
	t.Run("should return true for a royal flush hand", func(t *testing.T) {
		h := NewHand(
			card.TenOfHearts,
			card.JackOfHearts,
			card.QueenOfHearts,
			card.KingOfHearts,
			card.AceOfHearts,
		)
		assert.True(t, h.IsRoyalFlush())
	})

	t.Run("should return false for a straight flush hand", func(t *testing.T) {
		h := NewHand(
			card.NineOfHearts,
			card.TenOfHearts,
			card.JackOfHearts,
			card.QueenOfHearts,
			card.KingOfHearts,
		)
		assert.False(t, h.IsRoyalFlush())
	})
}

func TestIsFourOfAKind(t *testing.T) {
	t.Run("should return true for a four of a kind hand", func(t *testing.T) {
		h := NewHand(
			card.AceOfHearts,
			card.AceOfSpades,
			card.AceOfDiamonds,
			card.AceOfClubs,
			card.FiveOfHearts,
		)
		assert.True(t, h.IsFourOfAKind())
	})

	t.Run("should return false for a non-four of a kind hand", func(t *testing.T) {
		h := NewHand(
			card.AceOfHearts,
			card.AceOfSpades,
			card.AceOfDiamonds,
			card.KingOfClubs,
			card.FiveOfHearts,
		)
		assert.False(t, h.IsFourOfAKind())
	})
}

func TestIsFullHouse(t *testing.T) {
	t.Run("should return true for a full house hand", func(t *testing.T) {
		h := NewHand(
			card.AceOfHearts,
			card.AceOfSpades,
			card.AceOfDiamonds,
			card.KingOfClubs,
			card.KingOfHearts,
		)
		assert.True(t, h.IsFullHouse())
	})

	t.Run("should return false for a non-full house hand", func(t *testing.T) {
		h := NewHand(
			card.AceOfHearts,
			card.AceOfSpades,
			card.KingOfDiamonds,
			card.KingOfClubs,
			card.FiveOfHearts,
		)
		assert.False(t, h.IsFullHouse())
	})
}

func TestIsThreeOfAKind(t *testing.T) {
	t.Run("should return true for a three of a kind hand", func(t *testing.T) {
		h := NewHand(
			card.AceOfHearts,
			card.AceOfSpades,
			card.AceOfDiamonds,
			card.KingOfClubs,
			card.FiveOfHearts,
		)
		assert.True(t, h.IsThreeOfAKind())
	})

	t.Run("should return false for a non-three of a kind hand", func(t *testing.T) {
		h := NewHand(
			card.AceOfHearts,
			card.AceOfSpades,
			card.KingOfDiamonds,
			card.KingOfClubs,
			card.FiveOfHearts,
		)
		assert.False(t, h.IsThreeOfAKind())
	})

	t.Run("should return false for a full house hand", func(t *testing.T) {
		h := NewHand(
			card.AceOfHearts,
			card.AceOfSpades,
			card.AceOfDiamonds,
			card.KingOfClubs,
			card.KingOfHearts,
		)
		assert.False(t, h.IsThreeOfAKind())
	})
}

func TestIsTwoPair(t *testing.T) {
	t.Run("should return true for a two pair hand", func(t *testing.T) {
		h := NewHand(
			card.AceOfHearts,
			card.AceOfSpades,
			card.KingOfDiamonds,
			card.KingOfClubs,
			card.FiveOfHearts,
		)
		assert.True(t, h.IsTwoPair())
	})

	t.Run("should return false for a one pair hand", func(t *testing.T) {
		h := NewHand(
			card.AceOfHearts,
			card.AceOfSpades,
			card.QueenOfDiamonds,
			card.JackOfClubs,
			card.TenOfHearts,
		)
		assert.False(t, h.IsTwoPair())
	})

	t.Run("should return false for a full house hand", func(t *testing.T) {
		h := NewHand(
			card.AceOfHearts,
			card.AceOfSpades,
			card.AceOfDiamonds,
			card.KingOfClubs,
			card.KingOfHearts,
		)
		assert.False(t, h.IsTwoPair())
	})
}

func TestIsOnePair(t *testing.T) {
	t.Run("should return true for a one pair hand", func(t *testing.T) {
		h := NewHand(
			card.AceOfHearts,
			card.AceOfSpades,
			card.QueenOfDiamonds,
			card.JackOfClubs,
			card.TenOfHearts,
		)
		assert.True(t, h.IsOnePair())
	})

	t.Run("should return false for a two pair hand", func(t *testing.T) {
		h := NewHand(
			card.AceOfHearts,
			card.AceOfSpades,
			card.KingOfDiamonds,
			card.KingOfClubs,
			card.FiveOfHearts,
		)
		assert.False(t, h.IsOnePair())
	})
}
