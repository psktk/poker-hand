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
