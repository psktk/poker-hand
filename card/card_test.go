package card

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompare(t *testing.T) {
	t.Run("should return 1 if rank is higher", func(t *testing.T) {
		assert.Equal(t, 1, AceOfSpades.Compare(KingOfSpades))
	})

	t.Run("should return -1 if rank is lower", func(t *testing.T) {
		assert.Equal(t, -1, KingOfSpades.Compare(AceOfSpades))
	})

	t.Run("should return 0 if rank is same", func(t *testing.T) {
		assert.Equal(t, 0, KingOfSpades.Compare(KingOfHearts))
	})
}
