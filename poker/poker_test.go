package poker

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/psktk/poker-hand/card"
)

func TestIsFlush(t *testing.T) {
	t.Run("should return true for a straight flush hand", func(t *testing.T) {
		h := NewHand(
			card.New(card.Ace, card.Heart),
			card.New(card.Two, card.Heart),
			card.New(card.Three, card.Heart),
			card.New(card.Four, card.Heart),
			card.New(card.Five, card.Heart),
		)
		assert.True(t, h.IsFlush())
	})

	t.Run("should return true for a flush hand", func(t *testing.T) {
		h := NewHand(
			card.New(card.Two, card.Spade),
			card.New(card.Four, card.Spade),
			card.New(card.Six, card.Spade),
			card.New(card.Eight, card.Spade),
			card.New(card.Ten, card.Spade),
		)
		assert.True(t, h.IsFlush())
	})

	t.Run("should return false for a non-flush hand", func(t *testing.T) {
		h := NewHand(
			card.New(card.Two, card.Heart),
			card.New(card.Four, card.Heart),
			card.New(card.Six, card.Heart),
			card.New(card.Eight, card.Heart),
			card.New(card.Ten, card.Spade),
		)
		assert.False(t, h.IsFlush())
	})
}

func TestIsStraight(t *testing.T) {
	t.Run("should return true for a straight hand two to six", func(t *testing.T) {
		h := NewHand(
			card.New(card.Ace, card.Heart),
			card.New(card.Two, card.Spade),
			card.New(card.Three, card.Diamond),
			card.New(card.Four, card.Club),
			card.New(card.Five, card.Heart),
		)
		assert.True(t, h.IsStraight())
	})

	t.Run("should return true for a straight hand nine to king", func(t *testing.T) {
		h := NewHand(
			card.New(card.Nine, card.Heart),
			card.New(card.Ten, card.Spade),
			card.New(card.Jack, card.Diamond),
			card.New(card.Queen, card.Club),
			card.New(card.King, card.Heart),
		)
		assert.True(t, h.IsStraight())
	})

	t.Run("should return false for a hand with duplicate ranks", func(t *testing.T) {
		h := NewHand(
			card.New(card.Two, card.Heart),
			card.New(card.Three, card.Spade),
			card.New(card.Three, card.Diamond),
			card.New(card.Four, card.Club),
			card.New(card.Five, card.Heart),
		)
		assert.False(t, h.IsStraight())
	})

	t.Run("should return false for a non-straight hand", func(t *testing.T) {
		h := NewHand(
			card.New(card.Two, card.Heart),
			card.New(card.Four, card.Spade),
			card.New(card.Six, card.Diamond),
			card.New(card.Eight, card.Club),
			card.New(card.Ten, card.Heart),
		)
		assert.False(t, h.IsStraight())
	})

	t.Run("should return true for a straight hand ten to ace", func(t *testing.T) {
		h := NewHand(
			card.New(card.Ten, card.Heart),
			card.New(card.Jack, card.Spade),
			card.New(card.Queen, card.Diamond),
			card.New(card.King, card.Club),
			card.New(card.Ace, card.Heart),
		)
		assert.True(t, h.IsStraight())
	})

	t.Run("should return true for a straight hand ace to five", func(t *testing.T) {
		h := NewHand(
			card.New(card.Ace, card.Heart),
			card.New(card.Two, card.Spade),
			card.New(card.Three, card.Diamond),
			card.New(card.Four, card.Club),
			card.New(card.Five, card.Heart),
		)
		assert.True(t, h.IsStraight())
	})
}

func TestIsStraightFlush(t *testing.T) {
	t.Run("should return true for a straight flush hand", func(t *testing.T) {
		h := NewHand(
			card.New(card.Ace, card.Heart),
			card.New(card.Two, card.Heart),
			card.New(card.Three, card.Heart),
			card.New(card.Four, card.Heart),
			card.New(card.Five, card.Heart),
		)
		assert.True(t, h.IsStraightFlush())
	})

	t.Run("should return false for a flush hand", func(t *testing.T) {
		h := NewHand(
			card.New(card.Two, card.Spade),
			card.New(card.Four, card.Spade),
			card.New(card.Six, card.Spade),
			card.New(card.Eight, card.Spade),
			card.New(card.Ten, card.Spade),
		)
		assert.False(t, h.IsStraightFlush())
	})

	t.Run("should return false for a straight hand", func(t *testing.T) {
		h := NewHand(
			card.New(card.Two, card.Heart),
			card.New(card.Three, card.Spade),
			card.New(card.Four, card.Diamond),
			card.New(card.Five, card.Club),
			card.New(card.Six, card.Heart),
		)
		assert.False(t, h.IsStraightFlush())
	})

	t.Run("should return false for a non-straight flush hand", func(t *testing.T) {
		h := NewHand(
			card.New(card.Two, card.Heart),
			card.New(card.Four, card.Heart),
			card.New(card.Six, card.Heart),
			card.New(card.Eight, card.Heart),
			card.New(card.Ten, card.Spade),
		)
		assert.False(t, h.IsStraightFlush())
	})
}

func TestIsRoyalFlush(t *testing.T) {
	t.Run("should return true for a royal flush hand", func(t *testing.T) {
		h := NewHand(
			card.New(card.Ten, card.Heart),
			card.New(card.Jack, card.Heart),
			card.New(card.Queen, card.Heart),
			card.New(card.King, card.Heart),
			card.New(card.Ace, card.Heart),
		)
		assert.True(t, h.IsRoyalFlush())
	})

	t.Run("should return false for a straight flush hand", func(t *testing.T) {
		h := NewHand(
			card.New(card.Nine, card.Heart),
			card.New(card.Ten, card.Heart),
			card.New(card.Jack, card.Heart),
			card.New(card.Queen, card.Heart),
			card.New(card.King, card.Heart),
		)
		assert.False(t, h.IsRoyalFlush())
	})
}

func TestSort(t *testing.T) {
	t.Run("should sort the hand in ascending order by rank", func(t *testing.T) {
		h := NewHand(
			card.New(card.King, card.Heart),
			card.New(card.Three, card.Spade),
			card.New(card.Queen, card.Diamond),
			card.New(card.Five, card.Club),
			card.New(card.Ten, card.Heart),
		)
		h.Sort()
		expected := NewHand(
			card.New(card.Three, card.Spade),
			card.New(card.Five, card.Club),
			card.New(card.Ten, card.Heart),
			card.New(card.Queen, card.Diamond),
			card.New(card.King, card.Heart),
		)
		assert.Equal(t, expected, h)
	})

	t.Run("should sort a hand with duplicate ranks", func(t *testing.T) {
		h := NewHand(
			card.New(card.Five, card.Heart),
			card.New(card.Three, card.Spade),
			card.New(card.Five, card.Diamond),
			card.New(card.Two, card.Club),
			card.New(card.Three, card.Heart),
		)
		h.Sort()
		expected := NewHand(
			card.New(card.Two, card.Club),
			card.New(card.Three, card.Spade),
			card.New(card.Three, card.Heart),
			card.New(card.Five, card.Heart),
			card.New(card.Five, card.Diamond),
		)
		assert.Equal(t, expected, h)
	})

	t.Run("should sort a hand with ace as the highest rank", func(t *testing.T) {
		h := NewHand(
			card.New(card.Three, card.Heart),
			card.New(card.Ace, card.Spade),
			card.New(card.Ten, card.Diamond),
			card.New(card.Four, card.Club),
			card.New(card.King, card.Heart),
		)
		h.Sort()
		expected := NewHand(
			card.New(card.Three, card.Heart),
			card.New(card.Four, card.Club),
			card.New(card.Ten, card.Diamond),
			card.New(card.King, card.Heart),
			card.New(card.Ace, card.Spade),
		)
		assert.Equal(t, expected, h)
	})
}
