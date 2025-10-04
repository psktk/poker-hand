package poker

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/psktk/poker-hand/card"
	"github.com/psktk/poker-hand/rank"
	"github.com/psktk/poker-hand/suit"
)

func TestIsFlush(t *testing.T) {
	t.Run("should return true for a straight flush hand", func(t *testing.T) {
		h := NewHand(
			card.New(rank.Ace, suit.Heart),
			card.New(rank.Two, suit.Heart),
			card.New(rank.Three, suit.Heart),
			card.New(rank.Four, suit.Heart),
			card.New(rank.Five, suit.Heart),
		)
		assert.True(t, h.IsFlush())
	})

	t.Run("should return true for a flush hand", func(t *testing.T) {
		h := NewHand(
			card.New(rank.Two, suit.Spade),
			card.New(rank.Four, suit.Spade),
			card.New(rank.Six, suit.Spade),
			card.New(rank.Eight, suit.Spade),
			card.New(rank.Ten, suit.Spade),
		)
		assert.True(t, h.IsFlush())
	})

	t.Run("should return false for a non-flush hand", func(t *testing.T) {
		h := NewHand(
			card.New(rank.Two, suit.Heart),
			card.New(rank.Four, suit.Heart),
			card.New(rank.Six, suit.Heart),
			card.New(rank.Eight, suit.Heart),
			card.New(rank.Ten, suit.Spade),
		)
		assert.False(t, h.IsFlush())
	})
}

func TestIsStraight(t *testing.T) {
	t.Run("should return true for a straight hand two to six", func(t *testing.T) {
		h := NewHand(
			card.New(rank.Ace, suit.Heart),
			card.New(rank.Two, suit.Spade),
			card.New(rank.Three, suit.Diamond),
			card.New(rank.Four, suit.Club),
			card.New(rank.Five, suit.Heart),
		)
		assert.True(t, h.IsStraight())
	})

	t.Run("should return true for a straight hand nine to king", func(t *testing.T) {
		h := NewHand(
			card.New(rank.Nine, suit.Heart),
			card.New(rank.Ten, suit.Spade),
			card.New(rank.Jack, suit.Diamond),
			card.New(rank.Queen, suit.Club),
			card.New(rank.King, suit.Heart),
		)
		assert.True(t, h.IsStraight())
	})

	t.Run("should return false for a hand with duplicate ranks", func(t *testing.T) {
		h := NewHand(
			card.New(rank.Two, suit.Heart),
			card.New(rank.Three, suit.Spade),
			card.New(rank.Three, suit.Diamond),
			card.New(rank.Four, suit.Club),
			card.New(rank.Five, suit.Heart),
		)
		assert.False(t, h.IsStraight())
	})

	t.Run("should return false for a non-straight hand", func(t *testing.T) {
		h := NewHand(
			card.New(rank.Two, suit.Heart),
			card.New(rank.Four, suit.Spade),
			card.New(rank.Six, suit.Diamond),
			card.New(rank.Eight, suit.Club),
			card.New(rank.Ten, suit.Heart),
		)
		assert.False(t, h.IsStraight())
	})

	t.Run("should return true for a straight hand ten to ace", func(t *testing.T) {
		h := NewHand(
			card.New(rank.Ten, suit.Heart),
			card.New(rank.Jack, suit.Spade),
			card.New(rank.Queen, suit.Diamond),
			card.New(rank.King, suit.Club),
			card.New(rank.Ace, suit.Heart),
		)
		assert.True(t, h.IsStraight())
	})

	t.Run("should return true for a straight hand ace to five", func(t *testing.T) {
		h := NewHand(
			card.New(rank.Ace, suit.Heart),
			card.New(rank.Two, suit.Spade),
			card.New(rank.Three, suit.Diamond),
			card.New(rank.Four, suit.Club),
			card.New(rank.Five, suit.Heart),
		)
		assert.True(t, h.IsStraight())
	})
}

func TestIsStraightFlush(t *testing.T) {
	t.Run("should return true for a straight flush hand", func(t *testing.T) {
		h := NewHand(
			card.New(rank.Ace, suit.Heart),
			card.New(rank.Two, suit.Heart),
			card.New(rank.Three, suit.Heart),
			card.New(rank.Four, suit.Heart),
			card.New(rank.Five, suit.Heart),
		)
		assert.True(t, h.IsStraightFlush())
	})

	t.Run("should return false for a flush hand", func(t *testing.T) {
		h := NewHand(
			card.New(rank.Two, suit.Spade),
			card.New(rank.Four, suit.Spade),
			card.New(rank.Six, suit.Spade),
			card.New(rank.Eight, suit.Spade),
			card.New(rank.Ten, suit.Spade),
		)
		assert.False(t, h.IsStraightFlush())
	})

	t.Run("should return false for a straight hand", func(t *testing.T) {
		h := NewHand(
			card.New(rank.Two, suit.Heart),
			card.New(rank.Three, suit.Spade),
			card.New(rank.Four, suit.Diamond),
			card.New(rank.Five, suit.Club),
			card.New(rank.Six, suit.Heart),
		)
		assert.False(t, h.IsStraightFlush())
	})

	t.Run("should return false for a non-straight flush hand", func(t *testing.T) {
		h := NewHand(
			card.New(rank.Two, suit.Heart),
			card.New(rank.Four, suit.Heart),
			card.New(rank.Six, suit.Heart),
			card.New(rank.Eight, suit.Heart),
			card.New(rank.Ten, suit.Spade),
		)
		assert.False(t, h.IsStraightFlush())
	})
}

func TestIsRoyalFlush(t *testing.T) {
	t.Run("should return true for a royal flush hand", func(t *testing.T) {
		h := NewHand(
			card.New(rank.Ten, suit.Heart),
			card.New(rank.Jack, suit.Heart),
			card.New(rank.Queen, suit.Heart),
			card.New(rank.King, suit.Heart),
			card.New(rank.Ace, suit.Heart),
		)
		assert.True(t, h.IsRoyalFlush())
	})

	t.Run("should return false for a straight flush hand", func(t *testing.T) {
		h := NewHand(
			card.New(rank.Nine, suit.Heart),
			card.New(rank.Ten, suit.Heart),
			card.New(rank.Jack, suit.Heart),
			card.New(rank.Queen, suit.Heart),
			card.New(rank.King, suit.Heart),
		)
		assert.False(t, h.IsRoyalFlush())
	})
}

func TestIsFourOfAKind(t *testing.T) {
	t.Run("should return true for a four of a kind hand", func(t *testing.T) {
		h := NewHand(
			card.New(rank.Ace, suit.Heart),
			card.New(rank.Ace, suit.Spade),
			card.New(rank.Ace, suit.Diamond),
			card.New(rank.Ace, suit.Club),
			card.New(rank.Five, suit.Heart),
		)
		assert.True(t, h.IsFourOfAKind())
	})

	t.Run("should return false for a non-four of a kind hand", func(t *testing.T) {
		h := NewHand(
			card.New(rank.Ace, suit.Heart),
			card.New(rank.Ace, suit.Spade),
			card.New(rank.Ace, suit.Diamond),
			card.New(rank.King, suit.Club),
			card.New(rank.Five, suit.Heart),
		)
		assert.False(t, h.IsFourOfAKind())
	})
}

func TestSort(t *testing.T) {
	t.Run("should sort the hand in ascending order by rank", func(t *testing.T) {
		h := NewHand(
			card.New(rank.King, suit.Heart),
			card.New(rank.Three, suit.Spade),
			card.New(rank.Queen, suit.Diamond),
			card.New(rank.Five, suit.Club),
			card.New(rank.Ten, suit.Heart),
		)
		h.Sort()
		expected := NewHand(
			card.New(rank.Three, suit.Spade),
			card.New(rank.Five, suit.Club),
			card.New(rank.Ten, suit.Heart),
			card.New(rank.Queen, suit.Diamond),
			card.New(rank.King, suit.Heart),
		)
		assert.Equal(t, expected, h)
	})

	t.Run("should sort a hand with duplicate ranks", func(t *testing.T) {
		h := NewHand(
			card.New(rank.Five, suit.Heart),
			card.New(rank.Three, suit.Spade),
			card.New(rank.Five, suit.Diamond),
			card.New(rank.Two, suit.Club),
			card.New(rank.Three, suit.Heart),
		)
		h.Sort()
		expected := NewHand(
			card.New(rank.Two, suit.Club),
			card.New(rank.Three, suit.Spade),
			card.New(rank.Three, suit.Heart),
			card.New(rank.Five, suit.Heart),
			card.New(rank.Five, suit.Diamond),
		)
		assert.Equal(t, expected, h)
	})

	t.Run("should sort a hand with ace as the highest rank", func(t *testing.T) {
		h := NewHand(
			card.New(rank.Three, suit.Heart),
			card.New(rank.Ace, suit.Spade),
			card.New(rank.Ten, suit.Diamond),
			card.New(rank.Four, suit.Club),
			card.New(rank.King, suit.Heart),
		)
		h.Sort()
		expected := NewHand(
			card.New(rank.Three, suit.Heart),
			card.New(rank.Four, suit.Club),
			card.New(rank.Ten, suit.Diamond),
			card.New(rank.King, suit.Heart),
			card.New(rank.Ace, suit.Spade),
		)
		assert.Equal(t, expected, h)
	})
}
