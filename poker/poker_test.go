package poker

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/psktk/poker-hand/card"
	"github.com/psktk/poker-hand/rank"
)

func TestIsFlush(t *testing.T) {
	t.Run("should return true for a straight flush hand", func(t *testing.T) {
		h := NewHand(
			card.New(rank.Ace, card.Heart),
			card.New(rank.Two, card.Heart),
			card.New(rank.Three, card.Heart),
			card.New(rank.Four, card.Heart),
			card.New(rank.Five, card.Heart),
		)
		assert.True(t, h.IsFlush())
	})

	t.Run("should return true for a flush hand", func(t *testing.T) {
		h := NewHand(
			card.New(rank.Two, card.Spade),
			card.New(rank.Four, card.Spade),
			card.New(rank.Six, card.Spade),
			card.New(rank.Eight, card.Spade),
			card.New(rank.Ten, card.Spade),
		)
		assert.True(t, h.IsFlush())
	})

	t.Run("should return false for a non-flush hand", func(t *testing.T) {
		h := NewHand(
			card.New(rank.Two, card.Heart),
			card.New(rank.Four, card.Heart),
			card.New(rank.Six, card.Heart),
			card.New(rank.Eight, card.Heart),
			card.New(rank.Ten, card.Spade),
		)
		assert.False(t, h.IsFlush())
	})
}

func TestIsStraight(t *testing.T) {
	t.Run("should return true for a straight hand two to six", func(t *testing.T) {
		h := NewHand(
			card.New(rank.Ace, card.Heart),
			card.New(rank.Two, card.Spade),
			card.New(rank.Three, card.Diamond),
			card.New(rank.Four, card.Club),
			card.New(rank.Five, card.Heart),
		)
		assert.True(t, h.IsStraight())
	})

	t.Run("should return true for a straight hand nine to king", func(t *testing.T) {
		h := NewHand(
			card.New(rank.Nine, card.Heart),
			card.New(rank.Ten, card.Spade),
			card.New(rank.Jack, card.Diamond),
			card.New(rank.Queen, card.Club),
			card.New(rank.King, card.Heart),
		)
		assert.True(t, h.IsStraight())
	})

	t.Run("should return false for a hand with duplicate ranks", func(t *testing.T) {
		h := NewHand(
			card.New(rank.Two, card.Heart),
			card.New(rank.Three, card.Spade),
			card.New(rank.Three, card.Diamond),
			card.New(rank.Four, card.Club),
			card.New(rank.Five, card.Heart),
		)
		assert.False(t, h.IsStraight())
	})

	t.Run("should return false for a non-straight hand", func(t *testing.T) {
		h := NewHand(
			card.New(rank.Two, card.Heart),
			card.New(rank.Four, card.Spade),
			card.New(rank.Six, card.Diamond),
			card.New(rank.Eight, card.Club),
			card.New(rank.Ten, card.Heart),
		)
		assert.False(t, h.IsStraight())
	})

	t.Run("should return true for a straight hand ten to ace", func(t *testing.T) {
		h := NewHand(
			card.New(rank.Ten, card.Heart),
			card.New(rank.Jack, card.Spade),
			card.New(rank.Queen, card.Diamond),
			card.New(rank.King, card.Club),
			card.New(rank.Ace, card.Heart),
		)
		assert.True(t, h.IsStraight())
	})

	t.Run("should return true for a straight hand ace to five", func(t *testing.T) {
		h := NewHand(
			card.New(rank.Ace, card.Heart),
			card.New(rank.Two, card.Spade),
			card.New(rank.Three, card.Diamond),
			card.New(rank.Four, card.Club),
			card.New(rank.Five, card.Heart),
		)
		assert.True(t, h.IsStraight())
	})
}

func TestIsStraightFlush(t *testing.T) {
	t.Run("should return true for a straight flush hand", func(t *testing.T) {
		h := NewHand(
			card.New(rank.Ace, card.Heart),
			card.New(rank.Two, card.Heart),
			card.New(rank.Three, card.Heart),
			card.New(rank.Four, card.Heart),
			card.New(rank.Five, card.Heart),
		)
		assert.True(t, h.IsStraightFlush())
	})

	t.Run("should return false for a flush hand", func(t *testing.T) {
		h := NewHand(
			card.New(rank.Two, card.Spade),
			card.New(rank.Four, card.Spade),
			card.New(rank.Six, card.Spade),
			card.New(rank.Eight, card.Spade),
			card.New(rank.Ten, card.Spade),
		)
		assert.False(t, h.IsStraightFlush())
	})

	t.Run("should return false for a straight hand", func(t *testing.T) {
		h := NewHand(
			card.New(rank.Two, card.Heart),
			card.New(rank.Three, card.Spade),
			card.New(rank.Four, card.Diamond),
			card.New(rank.Five, card.Club),
			card.New(rank.Six, card.Heart),
		)
		assert.False(t, h.IsStraightFlush())
	})

	t.Run("should return false for a non-straight flush hand", func(t *testing.T) {
		h := NewHand(
			card.New(rank.Two, card.Heart),
			card.New(rank.Four, card.Heart),
			card.New(rank.Six, card.Heart),
			card.New(rank.Eight, card.Heart),
			card.New(rank.Ten, card.Spade),
		)
		assert.False(t, h.IsStraightFlush())
	})
}

func TestIsRoyalFlush(t *testing.T) {
	t.Run("should return true for a royal flush hand", func(t *testing.T) {
		h := NewHand(
			card.New(rank.Ten, card.Heart),
			card.New(rank.Jack, card.Heart),
			card.New(rank.Queen, card.Heart),
			card.New(rank.King, card.Heart),
			card.New(rank.Ace, card.Heart),
		)
		assert.True(t, h.IsRoyalFlush())
	})

	t.Run("should return false for a straight flush hand", func(t *testing.T) {
		h := NewHand(
			card.New(rank.Nine, card.Heart),
			card.New(rank.Ten, card.Heart),
			card.New(rank.Jack, card.Heart),
			card.New(rank.Queen, card.Heart),
			card.New(rank.King, card.Heart),
		)
		assert.False(t, h.IsRoyalFlush())
	})
}

func TestSort(t *testing.T) {
	t.Run("should sort the hand in ascending order by rank", func(t *testing.T) {
		h := NewHand(
			card.New(rank.King, card.Heart),
			card.New(rank.Three, card.Spade),
			card.New(rank.Queen, card.Diamond),
			card.New(rank.Five, card.Club),
			card.New(rank.Ten, card.Heart),
		)
		h.Sort()
		expected := NewHand(
			card.New(rank.Three, card.Spade),
			card.New(rank.Five, card.Club),
			card.New(rank.Ten, card.Heart),
			card.New(rank.Queen, card.Diamond),
			card.New(rank.King, card.Heart),
		)
		assert.Equal(t, expected, h)
	})

	t.Run("should sort a hand with duplicate ranks", func(t *testing.T) {
		h := NewHand(
			card.New(rank.Five, card.Heart),
			card.New(rank.Three, card.Spade),
			card.New(rank.Five, card.Diamond),
			card.New(rank.Two, card.Club),
			card.New(rank.Three, card.Heart),
		)
		h.Sort()
		expected := NewHand(
			card.New(rank.Two, card.Club),
			card.New(rank.Three, card.Spade),
			card.New(rank.Three, card.Heart),
			card.New(rank.Five, card.Heart),
			card.New(rank.Five, card.Diamond),
		)
		assert.Equal(t, expected, h)
	})

	t.Run("should sort a hand with ace as the highest rank", func(t *testing.T) {
		h := NewHand(
			card.New(rank.Three, card.Heart),
			card.New(rank.Ace, card.Spade),
			card.New(rank.Ten, card.Diamond),
			card.New(rank.Four, card.Club),
			card.New(rank.King, card.Heart),
		)
		h.Sort()
		expected := NewHand(
			card.New(rank.Three, card.Heart),
			card.New(rank.Four, card.Club),
			card.New(rank.Ten, card.Diamond),
			card.New(rank.King, card.Heart),
			card.New(rank.Ace, card.Spade),
		)
		assert.Equal(t, expected, h)
	})
}
