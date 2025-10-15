package rank

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompare(t *testing.T) {
	t.Run("should return 1 if rank is higher", func(t *testing.T) {
		assert.Equal(t, 1, Ace.Compare(King))
	})

	t.Run("should return -1 if rank is lower", func(t *testing.T) {
		assert.Equal(t, -1, King.Compare(Ace))
	})

	t.Run("should return 0 if rank is equal", func(t *testing.T) {
		assert.Equal(t, 0, Queen.Compare(Queen))
	})
}
